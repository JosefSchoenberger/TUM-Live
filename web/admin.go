package web

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/TUM-Dev/gocast/dao"
	"github.com/TUM-Dev/gocast/model"
	"github.com/TUM-Dev/gocast/tools"
	"github.com/TUM-Dev/gocast/tools/tum"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// AdminPage serves all administration pages. todo: refactor into multiple methods
func (r mainRoutes) AdminPage(c *gin.Context) {
	foundContext, exists := c.Get("TUMLiveContext")
	if !exists {
		sentry.CaptureException(errors.New("context should exist but doesn't"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	tumLiveContext := foundContext.(tools.TUMLiveContext)
	if tumLiveContext.User == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	var users []model.User
	_ = r.UsersDao.GetAllAdminsAndLecturers(&users)
	courses, err := r.CoursesDao.GetAdministeredCoursesByUserId(context.Background(), tumLiveContext.User.ID, "", 0)
	if err != nil {
		logger.Error("couldn't query courses for user.", "err", err)
		courses = []model.Course{}
	}
	workers, err := r.WorkerDao.GetAllWorkers()
	if err != nil {
		sentry.CaptureException(err)
	}
	lectureHalls := r.LectureHallsDao.GetAllLectureHalls()
	indexData := NewIndexData()
	indexData.TUMLiveContext = tumLiveContext
	page := GetPageString(c.Request.URL.Path)
	var notifications []model.Notification
	var tokens []dao.AllTokensDto
	var infopages []model.InfoPage
	var serverNotifications []model.ServerNotification
	switch page {
	case "notifications":
		found, err := r.NotificationsDao.GetAllNotifications()
		if err != nil {
			logger.Error("couldn't query notifications", "err", err)
		} else {
			notifications = found
		}
	case "token":
		tokens, err = r.TokenDao.GetAllTokens(tumLiveContext.User)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error("couldn't query tokens", "err", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	case "info-pages":
		infopages, err = r.InfoPageDao.GetAll()
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			logger.Error("couldn't query texts", "err", err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
	case "serverStats":
		streams, err := r.StreamsDao.GetAllStreams()
		if err != nil {
			logger.Error("Can't get all streams", "err", err)
			sentry.CaptureException(err)
			streams = []model.Stream{}
		}
		indexData.TUMLiveContext.Course = &model.Course{
			Model:   gorm.Model{ID: 0},
			Streams: streams,
		}
	case "serverNotifications":
		if res, err := r.ServerNotificationDao.GetAllServerNotifications(); err == nil {
			serverNotifications = res
		} else {
			logger.Warn("could not get all server notifications", "err", err)
		}
	}
	semesters := r.CoursesDao.GetAvailableSemesters(c)
	y, t := tum.GetCurrentSemester()
	err = templateExecutor.ExecuteTemplate(c.Writer, "admin.gohtml",
		AdminPageData{
			Users:               users,
			Courses:             courses,
			IndexData:           indexData,
			LectureHalls:        lectureHalls,
			Page:                page,
			Workers:             WorkersData{Workers: workers, Token: tools.Cfg.WorkerToken},
			Semesters:           semesters,
			CurY:                y,
			CurT:                t,
			Tokens:              TokensData{Tokens: tokens, RtmpProxyURL: tools.Cfg.RtmpProxyURL, User: tumLiveContext.User},
			InfoPages:           infopages,
			ServerNotifications: serverNotifications,
			Notifications:       notifications,
		})
	if err != nil {
		logger.Error("Error executing template admin.gohtml", "err", err)
	}
}

func GetPageString(s string) string {
	switch s {
	case "":
		return "schedule"
	case "/admin/users":
		return "users"
	case "/admin/lectureHalls":
		return "lectureHalls"
	case "/admin/lectureHalls/new":
		return "createLectureHalls"
	case "/admin/workers":
		return "workers"
	case "/admin/create-course":
		return "createCourse"
	case "/admin/course-import":
		return "courseImport"
	case "/admin/audits":
		return "audits"
	case "/admin/maintenance":
		return "maintenance"
	case "/admin/notifications":
		return "notifications"
	case "/admin/token":
		return "token"
	case "/admin/infopages":
		return "info-pages"
	case "/admin/server-stats":
		return "serverStats"
	case "/admin/server-notifications":
		return "serverNotifications"
	default:
		return "schedule"
	}
}

type WorkersData struct {
	Workers []model.Worker
	Token   string
}

type TokensData struct {
	Tokens       []dao.AllTokensDto
	RtmpProxyURL string
	User         *model.User
}

func (r mainRoutes) LectureCutPage(c *gin.Context) {
	foundContext, exists := c.Get("TUMLiveContext")
	if !exists {
		sentry.CaptureException(errors.New("context should exist but doesn't"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	tumLiveContext := foundContext.(tools.TUMLiveContext)
	if err := templateExecutor.ExecuteTemplate(c.Writer, "lecture-cut.gohtml", tumLiveContext); err != nil {
		logger.Error("Error executing template lecture-cut.gohtml", "err", err)
	}
}

func (r mainRoutes) LectureUnitsPage(c *gin.Context) {
	foundContext, exists := c.Get("TUMLiveContext")
	if !exists {
		sentry.CaptureException(errors.New("context should exist but doesn't"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	tumLiveContext := foundContext.(tools.TUMLiveContext)
	indexData := NewIndexData()
	indexData.TUMLiveContext = tumLiveContext
	if err := templateExecutor.ExecuteTemplate(c.Writer, "lecture-units.gohtml", LectureUnitsPageData{
		IndexData: indexData,
		Lecture:   *tumLiveContext.Stream,
		Units:     tumLiveContext.Stream.Units,
	}); err != nil {
		sentry.CaptureException(err)
	}
}

func (r mainRoutes) LectureStatsPage(c *gin.Context) {
	foundContext, exists := c.Get("TUMLiveContext")
	if !exists {
		sentry.CaptureException(errors.New("context should exist but doesn't"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	tumLiveContext := foundContext.(tools.TUMLiveContext)
	indexData := NewIndexData()
	indexData.TUMLiveContext = tumLiveContext
	if err := templateExecutor.ExecuteTemplate(c.Writer, "lecture-stats.gohtml", LectureStatsPageData{
		IndexData: indexData,
		Lecture:   *tumLiveContext.Stream,
	}); err != nil {
		sentry.CaptureException(err)
	}
}

func (r mainRoutes) CourseStatsPage(c *gin.Context) {
	foundContext, exists := c.Get("TUMLiveContext")
	if !exists {
		sentry.CaptureException(errors.New("context should exist but doesn't"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	tumLiveContext := foundContext.(tools.TUMLiveContext)
	indexData := NewIndexData()
	indexData.TUMLiveContext = tumLiveContext
	courses, err := r.CoursesDao.GetAdministeredCoursesByUserId(context.Background(), tumLiveContext.User.ID, "", 0)
	if err != nil {
		logger.Error("couldn't query courses for user.", "err", err)
		courses = []model.Course{}
	}
	semesters := r.CoursesDao.GetAvailableSemesters(c)
	err = templateExecutor.ExecuteTemplate(c.Writer, "admin.gohtml", AdminPageData{
		IndexData: indexData,
		Courses:   courses,
		Page:      "stats",
		Semesters: semesters,
		CurY:      tumLiveContext.Course.Year,
		CurT:      tumLiveContext.Course.TeachingTerm,
	})
	if err != nil {
		logger.Error("Error getting available semesters", "err", err)
	}
}

func (r mainRoutes) EditCoursePage(c *gin.Context) {
	foundContext, exists := c.Get("TUMLiveContext")
	if !exists {
		sentry.CaptureException(errors.New("context should exist but doesn't"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	tumLiveContext := foundContext.(tools.TUMLiveContext)
	lectureHalls := r.LectureHallsDao.GetAllLectureHalls()
	err := r.CoursesDao.GetInvitedUsersForCourse(tumLiveContext.Course)
	if err != nil {
		logger.Error("Error getting invited users for course", "err", err)
	}
	y, t := tum.GetCurrentSemester()

	indexData := NewIndexData()
	indexData.TUMLiveContext = tumLiveContext
	indexData.CurrentYear = y
	indexData.CurrentTerm = t

	courses, err := r.CoursesDao.GetAdministeredCoursesByUserId(context.Background(), tumLiveContext.User.ID, "", 0)
	if err != nil {
		logger.Error("couldn't query courses for user.", "err", err)
		courses = []model.Course{}
	}
	semesters := r.CoursesDao.GetAvailableSemesters(c)
	for i := range tumLiveContext.Course.Streams {
		err := tools.SetSignedPlaylists(&tumLiveContext.Course.Streams[i], tumLiveContext.User, true)
		if err != nil {
			logger.Error("could not set signed playlist for admin page", "err", err)
		}
	}
	err = templateExecutor.ExecuteTemplate(c.Writer, "admin.gohtml", AdminPageData{
		IndexData: indexData,
		Courses:   courses,
		Page:      "course",
		Semesters: semesters,
		CurY:      tumLiveContext.Course.Year,
		CurT:      tumLiveContext.Course.TeachingTerm,
		EditCourseData: EditCourseData{
			IndexData:    indexData,
			Courses:      courses,
			IngestBase:   tools.Cfg.IngestBase,
			LectureHalls: lectureHalls,
		},
	})
	if err != nil {
		logger.Error("Error executing template admin.gohtml", "err", err)
	}
}

func (r mainRoutes) UpdateCourse(c *gin.Context) {
	foundContext, exists := c.Get("TUMLiveContext")
	if !exists {
		sentry.CaptureException(errors.New("context should exist but doesn't"))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	tumLiveContext := foundContext.(tools.TUMLiveContext)
	if c.PostForm("submit") == "Reload Students From TUMOnline" {
		tum.FindStudentsForCourses([]model.Course{*tumLiveContext.Course}, r.UsersDao)
		c.Redirect(http.StatusFound, fmt.Sprintf("/admin/course/%v", tumLiveContext.Course.ID))
		return
	} else if c.PostForm("submit") == "Reload Lectures From TUMOnline" {
		tum.GetEventsForCourses([]model.Course{*tumLiveContext.Course}, r.DaoWrapper)
		c.Redirect(http.StatusFound, fmt.Sprintf("/admin/course/%v", tumLiveContext.Course.ID))
		return
	}
	access := c.PostForm("access")
	if match, err := regexp.MatchString("(public|loggedin|enrolled|hidden)", access); err != nil || !match {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"msg": "bad course id"})
		return
	}
	enVOD := c.PostForm("enVOD") == "on"
	enDL := c.PostForm("enDL") == "on"
	enChat := c.PostForm("enChat") == "on"
	enChatAnon := c.PostForm("enChatAnon") == "on"
	enChatMod := c.PostForm("enChatMod") == "on"
	livePrivate := c.PostForm("livePrivate") == "on"
	vodPrivate := c.PostForm("vodPrivate") == "on"
	tumLiveContext.Course.Visibility = access
	tumLiveContext.Course.VODEnabled = enVOD
	tumLiveContext.Course.DownloadsEnabled = enDL
	tumLiveContext.Course.ChatEnabled = enChat
	tumLiveContext.Course.AnonymousChatEnabled = enChatAnon
	tumLiveContext.Course.ModeratedChatEnabled = enChatMod
	tumLiveContext.Course.LivePrivate = livePrivate
	tumLiveContext.Course.VodPrivate = vodPrivate
	r.CoursesDao.UpdateCourseMetadata(context.Background(), *tumLiveContext.Course)
	c.Redirect(http.StatusFound, fmt.Sprintf("/admin/course/%v", tumLiveContext.Course.ID))
}

type AdminPageData struct {
	IndexData           IndexData
	Users               []model.User
	Courses             []model.Course
	LectureHalls        []model.LectureHall
	Page                string
	Workers             WorkersData
	Semesters           []model.Semester
	CurY                int
	CurT                string
	EditCourseData      EditCourseData
	ServerNotifications []model.ServerNotification
	Tokens              TokensData
	InfoPages           []model.InfoPage
	Notifications       []model.Notification
}

func (apd AdminPageData) UsersAsJson() string {
	type relevantUserInfo struct {
		ID    uint   `json:"id"`
		Name  string `json:"name"`
		Role  uint   `json:"role"`
		Email string `json:"email"`
	}

	users := make([]relevantUserInfo, len(apd.Users))
	for i, user := range apd.Users {
		users[i] = relevantUserInfo{
			ID:    user.ID,
			Name:  user.GetPreferredName(),
			Role:  user.Role,
			Email: user.Email.String,
		}
	}
	jsonStr, _ := json.Marshal(users)
	return string(jsonStr)
}

type EditCourseData struct {
	IndexData    IndexData
	IngestBase   string
	LectureHalls []model.LectureHall
	Courses      []model.Course // administered courses of user
}

type LectureUnitsPageData struct {
	IndexData IndexData
	Lecture   model.Stream
	Units     []model.StreamUnit
}

type LectureStatsPageData struct {
	IndexData IndexData
	Lecture   model.Stream
}

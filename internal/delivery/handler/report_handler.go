package handler

import (
	"fmt"
	"net/http"
	"strconv"
	reportusecase "github.com/bishal05das/blog_app_1/internal/usecase/report"
	util "github.com/bishal05das/blog_app_1/utils"
)

type ReportHandler struct {
	postUC reportusecase.Report
	userUC reportusecase.Report
	jsonUC reportusecase.Formatter
	htmlUC reportusecase.Formatter
}

func NewReportHandler(postUC reportusecase.Report, userUC reportusecase.Report, jsonUC reportusecase.Formatter,
	htmlUC reportusecase.Formatter) *ReportHandler {
	return &ReportHandler{
		postUC: postUC,
		userUC: userUC,
		jsonUC: jsonUC,
		htmlUC: htmlUC,
	}
}

func (rh *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("error in id converting into int", err)
		return
	}
	reportType := r.PathValue("report_type")
	formatType := r.URL.Query().Get("format")

	var formatter reportusecase.Formatter

	if formatType == "json" {
		formatter = rh.jsonUC
	} else {
		formatter = rh.htmlUC
	}

	var report reportusecase.Report
	if reportType == "users" {
		report = rh.userUC
	} else if reportType == "posts" {
		report = rh.postUC
	} else {
		http.Error(w, "Unknown report type", http.StatusBadRequest)
		return
	}

	output := report.Generate(id,formatter)
	util.SendData(w, output, http.StatusOK)

}

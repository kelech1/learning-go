package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
	"runtime"

	"github.com/kelech1/learning-go/internal/models"
	"github.com/kelech1/learning-go/internal/nlp"
)

var templates *template.Template

func init() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	templates = template.Must(template.ParseGlob(filepath.Join(basepath, "../../web/templates/*.html")))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "layout.html", nil)
}

func AnalyzeHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	result := nlp.Analyze(text)

	analysis := models.Analysis{
		Text:          text,
		SentenceCount: result.SentenceCount,
		WordCount:     result.WordCount,
		Sentiment:     result.Sentiment,
	}

	templates.ExecuteTemplate(w, "result.html", analysis)
}

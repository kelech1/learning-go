package nlp

import (
	"strings"

	"github.com/kelech1/learning-go/internal/models"
	"github.com/neurosnap/sentences"
)

func Analyze(text string) models.Analysis {
	tokenizer := sentences.NewSentenceTokenizer(nil)
	sents := tokenizer.Tokenize(text)

	wordCount := len(strings.Fields(text))

	positiveWords := []string{"good", "great", "excellent", "happy", "positive"}
	negativeWords := []string{"bad", "awful", "terrible", "sad", "negative"}

	sentimentScore := 0
	for _, word := range strings.Fields(strings.ToLower(text)) {
		if contains(positiveWords, word) {
			sentimentScore++
		} else if contains(negativeWords, word) {
			sentimentScore--
		}
	}

	var sentiment string
	if sentimentScore > 0 {
		sentiment = "Positive"
	} else if sentimentScore < 0 {
		sentiment = "Negative"
	} else {
		sentiment = "Neutral"
	}

	return models.Analysis{
		SentenceCount: len(sents),
		WordCount:     wordCount,
		Sentiment:     sentiment,
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

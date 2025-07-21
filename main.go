package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"temp/template"
	"time"

	"github.com/a-h/templ"
)

const PORT = ":8080"

var courses = []template.Course{
	{
		Name: "prog2",
		QuizPages: []template.MCTask{
			{
				Question: "Was ist eine Variable in der Programmierung?",
				Awnsers: []template.Anwser{
					{Ok: true, Content: "Ein Speicherplatz für einen veränderbaren Wert."},
					{Ok: false, Content: "Ein fester Wert, der sich nie ändert."},
					{Ok: false, Content: "Eine Schleife, die endlos läuft."},
					{Ok: false, Content: "Ein Fehler im Programm."},
				},
			},
			{
				Question: "Welche Schleife gibt es in Go?",
				Awnsers: []template.Anwser{
					{Ok: true, Content: "for"},
					{Ok: false, Content: "while"},
					{Ok: false, Content: "foreach"},
					{Ok: false, Content: "loop"},
				},
			},
		},
		CodePages: []template.CodeTask{
			{
				Question: "Schreibe ein Java-Programm, das 'Hello, World!' ausgibt.",
			},
		},
		LernPages: []template.LernTask{
			{
				Title:  "Was ist eine Funktion?",
				Layout: template.LayoutImgWithText,
			},
		},
		GapTextPages: []template.GapTextTask{
			{
				Question: "In Java wird eine Funktion mit dem Schlüsselwort ___ definiert.",
				Paragraph: []template.Paragraph{
					{
						Text:    "In Go wird eine Funktion mit dem Schlüsselwort  ",
						Anwsers: []string{"func", "void", "def"},
					},
					{
						Text:    "definiert",
						Anwsers: nil,
					},
				},
			},
		},
		TextPages: []template.TextTask{
			{
				Question: "Erläutre den Unterschied von Methoden und Funktionen.",
			},
		},
		Pages: []template.CoursePageType{
			template.PageLern,
			template.PageCode,
			template.PageGapText,
			template.PageQuiz,
			template.PageText,
		},
	},
	{
		Name: "mathematik",
		QuizPages: []template.MCTask{
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Eine Funktion, die nur positive Werte annimmt."},
					{Ok: true, Content: "Eine Funktion, deren Ableitung gleich der Funktion selbst ist."},
					{Ok: false, Content: "Eine Funktion, die nur diskrete Werte annimmt."},
					{Ok: false, Content: "Eine Funktion, die immer linear ist."},
				},
				Question: "Was ist eine Exponentialfunktion?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Die Menge aller reellen Zahlen."},
					{Ok: true, Content: "Die Menge aller natürlichen Zahlen ohne Null."},
					{Ok: false, Content: "Die Menge aller komplexen Zahlen."},
					{Ok: false, Content: "Die Menge aller rationalen Zahlen."},
				},
				Question: "Was ist die Definition der Menge der natürlichen Zahlen N+?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Eine Reihe von Operationen, die nicht terminieren."},
					{Ok: true, Content: "Eine endliche Folge von eindeutig definierten Anweisungen zur Lösung eines Problems."},
					{Ok: false, Content: "Ein Programmcode ohne Kommentare."},
					{Ok: false, Content: "Eine zufällige Anordnung von Befehlen."},
				},
				Question: "Was ist ein Algorithmus?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Addition und Subtraktion"},
					{Ok: true, Content: "Multiplikation und Addition"},
					{Ok: false, Content: "Division und Subtraktion"},
					{Ok: false, Content: "Potenzierung und Wurzelziehen"},
				},
				Question: "Welche Operationen sind für die Bildung eines Vektorraums zwingend erforderlich?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Eine Menge, die nur eine Teilmenge ist."},
					{Ok: true, Content: "Eine Menge, die alle ihre Häufungspunkte enthält."},
					{Ok: false, Content: "Eine Menge, die unendlich viele Elemente hat."},
					{Ok: false, Content: "Eine Menge, die keine inneren Punkte besitzt."},
				},
				Question: "Was bedeutet es, dass eine Menge 'abgeschlossen' ist?",
			},
		},
		LernPages: []template.LernTask{
			{
				Title:  "Einführung in Mathematik I",
				Layout: template.LayoutTwoImagesSomeText,
			},
		},
		CodePages: nil,
		Pages: []template.CoursePageType{
			template.PageLern,
			template.PageQuiz,
			template.PageQuiz,
			template.PageQuiz,
			template.PageQuiz,
			template.PageQuiz,
		},
	},
	{
		Name: "mas",
		QuizPages: []template.MCTask{
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Unified Processing Model"},
					{Ok: true, Content: "Unified Modeling Language"},
					{Ok: false, Content: "Universal Machine Language"},
					{Ok: false, Content: "Unique Management Logic"},
				},
				Question: "Wofür steht die Abkürzung UML in der Systemmodellierung?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Sequenzdiagramm"},
					{Ok: false, Content: "Klassendiagramm"},
					{Ok: true, Content: "Anwendungsfalldiagramm"},
					{Ok: false, Content: "Aktivitätsdiagramm"},
				},
				Question: "Welches UML-Diagramm wird verwendet, um die Interaktion von Benutzern mit einem System darzustellen?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Eine Methode zur Programmierung von Benutzerschnittstellen."},
					{Ok: true, Content: "Ein Prozess, bei dem die Systemfunktionalität in kleinere, wiederverwendbare Einheiten unterteilt wird."},
					{Ok: false, Content: "Ein Werkzeug zur Versionskontrolle."},
					{Ok: false, Content: "Eine Art von Datenbankmodell."},
				},
				Question: "Was versteht man unter 'Modularisierung' in der Systementwicklung?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Ein Diagramm zur Darstellung von Datenflüssen."},
					{Ok: true, Content: "Ein Diagramm zur Beschreibung des Verhaltens von Objekten im Laufe der Zeit, basierend auf Zuständen und Transitionen."},
					{Ok: false, Content: "Ein Diagramm zur Modellierung von Hardwarekomponenten."},
					{Ok: false, Content: "Ein Diagramm zur Darstellung der Klassenhierarchie."},
				},
				Question: "Wozu dient ein Zustandsdiagramm in der UML?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Ein Framework für die Softwareentwicklung."},
					{Ok: true, Content: "Ein iterativer und inkrementeller Ansatz zur Softwareentwicklung, bei dem Anforderungen und Lösungen durch die Zusammenarbeit von selbstorganisierenden, funktionsübergreifenden Teams entstehen."},
					{Ok: false, Content: "Eine starre, sequenzielle Methode zur Projektplanung."},
					{Ok: false, Content: "Ein Modell zur Kostenkalkulation von Softwareprojekten."},
				},
				Question: "Was ist 'Agile Softwareentwicklung'?",
			},
		},
		CodePages: nil,
		TextPages: []template.TextTask{
			{
				Question: "Was versteht man unter einem Anwendungssystem, und welche Rolle spielt die Modellierung in dessen Entwicklung?",
			},
		},
		LernPages: []template.LernTask{
			{
				Title:  "Einführung in Modellierung von Anwendungssystemen",
				Layout: template.LayoutOnlyText,
			},
		},
		Pages: []template.CoursePageType{
			template.PageLern,
			template.PageText,
			template.PageQuiz,
			template.PageQuiz,
			template.PageQuiz,
			template.PageQuiz,
		},
	},
	{
		Name: "bean",
		QuizPages: []template.MCTask{
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Employee Resource Planning"},
					{Ok: true, Content: "Enterprise Resource Planning"},
					{Ok: false, Content: "External Relations Protocol"},
					{Ok: false, Content: "Environmental Risk Protection"},
				},
				Question: "Wofür steht die Abkürzung ERP im Kontext von betrieblichen Anwendungen?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Ein System zur Verwaltung von E-Mail-Kampagnen."},
					{Ok: true, Content: "Ein integriertes Softwaresystem, das die wichtigsten Geschäftsprozesse eines Unternehmens wie Finanzwesen, Personalwesen, Produktion und Vertrieb verwaltet und automatisiert."},
					{Ok: false, Content: "Ein Tool zur Netzwerküberwachung."},
					{Ok: false, Content: "Eine Anwendung zur Erstellung von Präsentationen."},
				},
				Question: "Was ist ein ERP-System?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Customer Reporting Management"},
					{Ok: true, Content: "Customer Relationship Management"},
					{Ok: false, Content: "Client Retention Metrics"},
					{Ok: false, Content: "Content Resource Management"},
				},
				Question: "Wofür steht die Abkürzung CRM?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Ein System zur Verwaltung von Lieferantenbeziehungen."},
					{Ok: true, Content: "Ein System zur Verwaltung der Interaktionen eines Unternehmens mit aktuellen und potenziellen Kunden."},
					{Ok: false, Content: "Ein System zur Bestandsverwaltung im Lager."},
					{Ok: false, Content: "Eine Software zur Finanzanalyse."},
				},
				Question: "Was ist ein CRM-System?",
			},
			{
				Awnsers: []template.Anwser{
					{Ok: false, Content: "Eine Anwendung zur Verwaltung von Mitarbeiterdaten."},
					{Ok: true, Content: "Ein System, das die gesamte Lieferkette eines Unternehmens optimiert, von der Beschaffung der Rohstoffe über die Produktion bis zur Auslieferung an den Kunden."},
					{Ok: false, Content: "Eine Software zur Planung von Projektbudgets."},
					{Ok: false, Content: "Ein Tool zur Erstellung von Geschäftsdokumenten."},
				},
				Question: "Was versteht man unter 'Supply Chain Management' (SCM) Software?",
			},
		},
		LernPages: []template.LernTask{
			{
				Title:  "Einführung in Geschäftsprozesse und betriebliche Anwendungen",
				Layout: template.LayoutImgWithText,
			},
		},
		CodePages: nil,
		Pages: []template.CoursePageType{
			template.PageLern,
			template.PageQuiz,
			template.PageQuiz,
			template.PageQuiz,
			template.PageQuiz,
			template.PageQuiz,
		},
	},
}

var currentQuiz string

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)
		log.Println(wrapped.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}

func getCouse(c *[]template.Course, name string) (*template.Course, bool) {
	for i := range *c {
		if (*c)[i].Name == name {
			return &(*c)[i], true
		}
	}

	return nil, false
}

func main() {
	router := http.NewServeMux()

	router.Handle("/", http.FileServer(http.Dir("include_dir")))
	router.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "")
	})

	router.HandleFunc("/ubung/{name}/{id}", func(w http.ResponseWriter, r *http.Request) {

		currentQuiz = r.PathValue("name")
		id_str := r.PathValue("id")

		c, ok := getCouse(&courses, currentQuiz)

		if !ok {
			w.Header().Add("Hx-Redirect", "/dashboard.html")
			return
		}

		id, _ := strconv.Atoi(id_str)
		c.Progress = id
		log.Println(id)

		log.Println(courses)

		w.Header().Add("Hx-Redirect", "/ubung.html")
	})

	router.HandleFunc("/quiz", func(w http.ResponseWriter, r *http.Request) {
		c, ok := getCouse(&courses, currentQuiz)
		if !ok {
			w.Header().Add("Hx-Redirect", "/dashboard.html")
			return
		}

		template.RenderPage(c).Render(context.TODO(), w)
	})

	log.Println("Starting Server at :8080")
	router.Handle("/ai-settings", templ.Handler(template.AISetting()))
	router.Handle("/header", templ.Handler(template.Header()))
	router.HandleFunc("/feedback", func(w http.ResponseWriter, r *http.Request) {
		if currentQuiz == "" {
			w.Header().Add("Hx-Redirect", "/dashboard.html")
			return
		}

		if currentQuiz == "intro" {
			template.KalibrungFeedback().Render(context.TODO(), w)
			return
		}

		c, ok := getCouse(&courses, currentQuiz)

		if !ok {
			w.Header().Add("Hx-Redirect", "/dashboard.html")
			return
		}

		template.Feedback(currentQuiz, len(c.Pages)).Render(context.TODO(), w)
	})

	log.Fatal(http.ListenAndServe(PORT, Logging(router)))

}

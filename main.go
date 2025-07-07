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

var quizId int = 0
var quiz map[string][]template.QuizMC = map[string][]template.QuizMC{
	"intro": {
		{
			Awnsers: []template.Anwser{
				{Ok: false, Content: "f'(x) = x"},
				{Ok: true, Content: "f'(x) = 2x"},
				{Ok: false, Content: "f'(x) = x²"},
				{Ok: false, Content: "f'(x) = 2"},
			},
			Question: "Was ist die Ableitung von f(x) = x²?",
		},
		{
			Awnsers: []template.Anwser{
				{Ok: false, Content: "O(1)"},
				{Ok: false, Content: "O(n)"},
				{Ok: false, Content: "O(log n)"},
				{Ok: false, Content: "O(n²)"},
			},
			Question: "Was ist die Zeitkomplexität einer linearen Suche?",
		},
	},
	"mathematik": {
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
	"mas": {
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
	"bean": {
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

func main() {
	router := http.NewServeMux()

	router.Handle("/", http.FileServer(http.Dir("include_dir")))
	router.HandleFunc("/del", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "")
	})

	router.HandleFunc("/ubung/{name}/{id}", func(w http.ResponseWriter, r *http.Request) {

		currentQuiz = r.PathValue("name")
		id_str := r.PathValue("id")
		quizId, _ = strconv.Atoi(id_str)
		w.Header().Add("Hx-Redirect", "/ubung.html")
	})

	router.HandleFunc("/quiz", func(w http.ResponseWriter, r *http.Request) {
		q, ok := quiz[currentQuiz]
		if !ok {
			w.Header().Add("Hx-Redirect", "/dashboard.html")
			return
		}

		template.MC(q, quizId, currentQuiz).Render(context.TODO(), w)
	})

	log.Println("Starting Server at :8080")
	router.Handle("/tmpl", templ.Handler(template.Test()))
	router.HandleFunc("/feedback", func(w http.ResponseWriter, r *http.Request) {
		if currentQuiz == "" {
			w.Header().Add("Hx-Redirect", "/dashboard.html")
			return
		}
		template.Feedback(currentQuiz, len(quiz[currentQuiz])).Render(context.TODO(), w)
	})

	log.Fatal(http.ListenAndServe(PORT, Logging(router)))

}

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"sync"
)

type MenuItem struct {
	ID    uint
	Name  string
	Price float64
}

type OrderItem struct {
	Name     string
	Quantity uint
	Price    float64
	Total    float64
}

var menu = []MenuItem{
	{1, "Adrakh Chai", 20},
	{2, "Filter Coffee", 25},
	{3, "Chhas", 35.50},
	{4, "Lassi", 30},
	{5, "Mango Lassi", 60},
	{6, "Kadi Chaawal", 50},
	{7, "Chhole Bhature", 45},
	{8, "Khasta Kachori", 30},
	{9, "Raj Kachori", 30},
	{10, "Veg. Sandwich", 20},
	{11, "Veg. Masala Maggi", 60.00},
	{12, "Samosa", 20},
	{13, "Cream Roll", 15},
}

var (
	currentOrder []OrderItem
	orderMutex   sync.Mutex
	tmpl         = template.Must(template.ParseGlob("templates/*.html"))
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/menu", menuHandler)
	http.HandleFunc("/order", orderHandler)
	http.HandleFunc("/view-order", viewOrderHandler)
	http.HandleFunc("/bill", billHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func menuHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.ExecuteTemplate(w, "menu.html", menu); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func orderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/menu", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	var orderedItems []OrderItem
	for _, item := range menu {
		qtyStr := r.FormValue(fmt.Sprintf("item_%d", item.ID))
		qty, err := strconv.Atoi(qtyStr)
		if err == nil && qty > 0 {
			total := float64(qty) * item.Price
			orderedItems = append(orderedItems, OrderItem{
				Name:     item.Name,
				Quantity: uint(qty),
				Price:    item.Price,
				Total:    total,
			})
		}
	}

	orderMutex.Lock()
	currentOrder = orderedItems
	orderMutex.Unlock()

	http.Redirect(w, r, "/view-order", http.StatusSeeOther)
}

func viewOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderMutex.Lock()
	orderedItems := currentOrder
	orderMutex.Unlock()

	subtotal := 0.0
	for _, item := range orderedItems {
		subtotal += item.Total
	}
	gst := subtotal * 0.05
	total := subtotal + gst

	data := struct {
		Items    []OrderItem
		Subtotal float64
		GST      float64
		Total    float64
	}{
		Items:    orderedItems,
		Subtotal: subtotal,
		GST:      gst,
		Total:    total,
	}

	if err := tmpl.ExecuteTemplate(w, "order_summary.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func billHandler(w http.ResponseWriter, r *http.Request) {
	orderMutex.Lock()
	orderedItems := currentOrder
	orderMutex.Unlock()

	subtotal := 0.0
	for _, item := range orderedItems {
		subtotal += item.Total
	}
	gst := subtotal * 0.05
	total := subtotal + gst

	data := struct {
		Items    []OrderItem
		Subtotal float64
		GST      float64
		Total    float64
	}{
		Items:    orderedItems,
		Subtotal: subtotal,
		GST:      gst,
		Total:    total,
	}

	if err := tmpl.ExecuteTemplate(w, "bill.html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

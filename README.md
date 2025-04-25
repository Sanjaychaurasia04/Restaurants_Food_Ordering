# 🍽️ Restaurants Food Ordering System

A simple and interactive web-based food ordering system built with **Go**, **HTML**, and **CSS**. It allows users to view a menu, place food orders, and generate a final bill with GST calculation.

---

## 🚀 Features

- Browse a list of Indian food & beverages
- Place multiple item orders with quantity
- Dynamic order summary and total calculation
- Bill generation with 5% GST
- Thread-safe order handling using `sync.Mutex`
- Clean and responsive frontend using HTML & CSS

---

## 🛠️ Built With

- [Go (Golang)](https://golang.org/)
- HTML & CSS
- Go's `html/template` package

---

## 📁 Project Structure


restaurants_food_ordering/ ├── main.go ├── templates/ │ ├── index.html │ ├── menu.html │ ├── order_summary.html │ └── bill.html └── static/ └── style.css 



---

## 🧑‍🍳 How to Run

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-username/restaurants_food_ordering.git
   cd restaurants_food_ordering

2 **Run the application**

go run main.go

3 **Open in browser**

http://localhost:8080


# 🌐 Routes
/ → Home page

/menu → Menu display
/order → Order submission (POST)

/view-order → Order summary with GST

/bill → Final bill page
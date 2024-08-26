# **Customer Support Router Microservice**

## **Overview**

The **Customer Support Router Microservice** enhances customer support operations by routing support tickets to the appropriate agents based on predefined rules. Built using the **Golang Gin web framework**, this project demonstrates my ability to rapidly develop scalable and reliable backend services.

## **Features**

- **RESTful API** for ticket creation, retrieval, and routing.
- **High-Performance Routing** logic for efficient support ticket management.
- **Scalability** designed for cloud deployment and future AI enhancements.

## **Getting Started**

### **Prerequisites**

- **Golang** 1.23.0 or higher

### **Installation**

Clone the repository and install dependencies:

```bash
git clone https://github.com/foreverLoveWisdom/customer_support_microservice.git
cd customer_support_microservice
go mod tidy
```

### **Running the Application**

Start the server:

```bash
go run main.go
```

Access the service at `http://localhost:8080`.

## **API Endpoints**

- **GET** `/ping`: Health check.
- **POST** `/tickets`: Create a new support ticket.
- **GET** `/tickets`: Retrieve all support tickets.

## **Testing the Endpoints**

Test using `curl`:

- **Health Check**:

  ```bash
  curl http://localhost:8080/ping
  ```

- **Create a Ticket**:

  ```bash
  curl -X POST http://localhost:8080/tickets -H "Content-Type: application/json" -d '{"id":1,"subject":"Issue with order","description":"Details of the issue"}'
  ```

- **Get All Tickets**:

  ```bash
  curl http://localhost:8080/tickets
  ```

## **Future Enhancements**

- **AI-Driven Ticket Prioritization**
- **Real-Time Analytics Dashboard**
- **Advanced Routing Logic**

## **Why This Project?**

This project aligns with the needs of companies like Grab, demonstrating rapid learning and practical application of scalable solutions to real-world problems.

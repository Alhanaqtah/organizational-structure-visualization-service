# Organizational Structure Visualization Service

This project provides a web application to visualize an organization's structure and find employees by position, role, or other attributes.

---

## **Setup Instructions**

### **1. Clone the Repository**
```bash
git clone https://github.com/Alhanaqtah/organizational-structure-visualization-service
cd organizational-structure-visualization-service
```

---

### **2. Frontend Setup**
Navigate to the frontend directory, install dependencies, and start the development server:
```bash
cd frontend
npm install
npm run dev
```

---

### **3. Backend Setup**

#### **a. Prepare `.env` File**
Before starting the backend, create a `.env` file in the `backend` directory with the following content:

```plaintext
ENV=local

POSTGRES_URL=postgres://postgres:postgres@localhost:2345/postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres

HTTP_SERVER_ADDRESS=localhost:8080
HTTP_SERVER_IDLE=4
```

#### **b. Start PostgreSQL Using Docker**
Make sure you have Docker installed and running. Use the following command to start a PostgreSQL container:
```bash
docker run --name org-structure-db -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -p 2345:5432 -d postgres
```

#### **c. Start the Backend Server**
Navigate to the backend directory and run the service:
```bash
cd ../backend
go run ./cmd/main
```

---

## **Technology Stack**
- **Frontend**: React, JavaScript/TypeScript
- **Backend**: Go
- **Database**: PostgreSQL (Dockerized)

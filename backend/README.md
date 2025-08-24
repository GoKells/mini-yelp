# Backend System PRD: Café Management System

## Overview
This document outlines the requirements for the backend system of the Café Management System, built with Go. The backend serves as the core engine handling business logic, data persistence, API endpoints, and real-time communication for the entire application.

## Technology Stack
- **Language**: Go 1.21+
- **Framework**: Fiber Web Framework
- **Database**: PostgreSQL with Prisma ORM
- **Cache**: Redis
- **Real-time Communication**: WebSockets (Gorilla WebSocket)
- **API Documentation**: Swagger/OpenAPI
- **Authentication**: JWT with refresh tokens
- **Deployment**: Docker containers, Kubernetes-ready

## Architecture
```
backend/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── api/                     # HTTP handlers & routes
│   ├── service/                 # Business logic layer
│   ├── repository/              # Data access layer (Prisma)
│   ├── models/                  # Data structures
│   ├── middleware/              # Custom middleware
│   └── utils/                   # Helper functions
├── pkg/
│   ├── auth/                    # Authentication utilities
│   ├── database/                # Database connection & setup
│   └── websocket/               # WebSocket management
├── config/                      # Configuration files
├── migrations/                  # Database migrations
├── scripts/                     # Deployment & utility scripts
└── tests/                       # Integration & unit tests
```

## Core Modules

### 1. Authentication Service
**Endpoints:**
- `POST /api/auth/login` - Staff login
- `POST /api/auth/refresh` - Refresh JWT token
- `POST /api/auth/logout` - Invalidate token

**Requirements:**
- JWT-based authentication with 15-minute expiry
- Refresh tokens with 7-day expiry
- Role-based access control (Admin, Manager, Barista)
- Secure password hashing (bcrypt)

### 2. Order Management Module
**Endpoints:**
- `POST /api/orders` - Create new order
- `GET /api/orders` - List orders with filtering
- `GET /api/orders/{id}` - Get order details
- `PUT /api/orders/{id}/status` - Update order status
- `DELETE /api/orders/{id}` - Cancel order

**Real-time Events:**
- `order_created` - New order notification
- `order_updated` - Status change notification
- `order_completed` - Order completion notification

### 3. Menu Management Module
**Endpoints:**
- `GET /api/menu` - Get current menu
- `POST /api/menu/items` - Add menu item (Admin only)
- `PUT /api/menu/items/{id}` - Update menu item
- `DELETE /api/menu/items/{id}` - Remove menu item
- `PUT /api/menu/items/{id}/stock` - Update stock status

### 4. Inventory Management Module
**Endpoints:**
- `GET /api/inventory` - Get current inventory
- `POST /api/inventory` - Add inventory item
- `PUT /api/inventory/{id}` - Update inventory item
- `GET /api/inventory/low-stock` - Get low stock alerts
- `POST /api/inventory/adjustment` - Manual stock adjustment

### 5. Reporting Module
**Endpoints:**
- `GET /api/reports/sales/daily` - Daily sales report
- `GET /api/reports/sales/weekly` - Weekly sales report
- `GET /api/reports/products/top` - Top selling products
- `GET /api/reports/performance` - Staff performance metrics

## Database Schema (Key Tables)
```sql
-- Orders
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    type VARCHAR(20), -- dine-in, takeaway, online
    status VARCHAR(20), -- pending, in-progress, completed, cancelled
    total_amount DECIMAL(10,2),
    created_at TIMESTAMP,
    completed_at TIMESTAMP
);

-- Order Items
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id INTEGER REFERENCES orders(id),
    menu_item_id INTEGER REFERENCES menu_items(id),
    quantity INTEGER,
    special_instructions TEXT
);

-- Menu Items
CREATE TABLE menu_items (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    category VARCHAR(100),
    price DECIMAL(10,2),
    in_stock BOOLEAN DEFAULT true
);

-- Inventory
CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    current_stock DECIMAL(10,2),
    unit VARCHAR(20),
    low_stock_threshold DECIMAL(10,2)
);
```

## API Specifications

### Request/Response Examples

**Create Order:**
```http
POST /api/orders
Authorization: Bearer <token>
Content-Type: application/json

{
  "type": "dine-in",
  "items": [
    {
      "menu_item_id": 1,
      "quantity": 2,
      "modifiers": ["skim-milk", "extra-shot"]
    }
  ]
}

Response: 201 Created
{
  "id": 123,
  "status": "pending",
  "estimated_time": 10,
  "total_amount": 9.50
}
```

**Get Sales Report:**
```http
GET /api/reports/sales/daily?date=2023-10-26
Authorization: Bearer <token>

Response: 200 OK
{
  "date": "2023-10-26",
  "total_sales": 1250.75,
  "order_count": 143,
  "payment_methods": {
    "cash": 450.25,
    "card": 625.30,
    "mobile": 175.20
  }
}
```

## Performance Requirements
- API response time < 100ms for 95% of requests
- Support 100+ concurrent WebSocket connections
- Handle 50+ orders per minute during peak hours
- Database queries optimized with proper indexing

## Security Requirements
- All endpoints HTTPS only
- SQL injection prevention (parameterized queries)
- XSS protection
- Rate limiting on authentication endpoints
- Regular security audits and dependency updates

## Deployment & DevOps
- Docker containerization
- Kubernetes deployment manifests
- GitHub Actions CI/CD pipeline
- Environment-based configuration
- Health check endpoints (`/health`, `/ready`)
- Prometheus metrics endpoint (`/metrics`)

## Testing Strategy
- Unit tests for all service layer functions
- Integration tests for API endpoints
- Load testing for critical pathways
- Database migration testing
- WebSocket connection testing

## Monitoring & Logging
- Structured JSON logging
- Error tracking with Sentry
- Performance monitoring with Prometheus
- Audit logs for sensitive operations

## Milestones & Timeline

### Phase 1 (2 weeks): Core Foundation
- Project structure setup
- Database configuration with Prisma
- Basic authentication system
- Health check endpoints

### Phase 2 (3 weeks): Order Management
- Order creation API
- Real-time order updates via WebSockets
- Basic menu management
- Unit test coverage > 70%

### Phase 3 (2 weeks): Reporting & Inventory
- Sales reporting endpoints
- Inventory management system
- Low stock alerts
- API documentation

### Phase 4 (1 week): Polish & Deployment
- Performance optimization
- Security hardening
- Docker deployment setup
- Load testing

## Dependencies
- External: PostgreSQL, Redis
- Payment: Stripe API integration
- Optional: Third-party delivery service APIs

## Success Metrics
- 99.9% API availability
- < 100ms response time for core endpoints
- Zero critical security vulnerabilities
- 100% test coverage for critical pathways

This PRD serves as the foundation for the backend development of the Café Management System. All implementations should reference this document to ensure consistency with project requirements.
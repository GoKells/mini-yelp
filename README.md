# Café Management System

A modern, full-stack café management system built with React (frontend) and Go (backend). This system streamlines café operations with order management, inventory tracking, and customer engagement features.

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-1.21%2B-blue)
![React Version](https://img.shields.io/badge/react-18%2B-blue)

## 🚀 Features

### Core Features (Implemented)
- **POS System**: Process dine-in, takeaway, and online orders
- **Menu Management**: Dynamic menu with categories and modifiers
- **Payment Processing**: Multiple payment methods (cash, card, mobile wallets)
- **Real-time Updates**: Live order synchronization between POS and kitchen
- **Basic Reporting**: Sales summaries and analytics

### Upcoming Features
- Inventory management with low-stock alerts
- Customer loyalty program
- Mobile app for customers
- Multi-branch support

## 🛠️ Tech Stack

### Frontend
- **React 18** with TypeScript
- **Next.js** for server-side rendering
- **Tailwind CSS** for styling
- **React Query** for state management
- **Socket.io Client** for real-time updates

### Backend
- **Go 1.21+** (Gin framework)
- **PostgreSQL** with Prisma ORM
- **Redis** for caching and real-time features
- **Socket.io** for WebSocket connections
- **Stripe** for payment processing

## 📦 Installation

### Prerequisites
- Go 1.21+
- Node.js 18+
- PostgreSQL 12+
- Redis 6+

### Backend Setup

1. Clone the repository:
```bash
git clone https://github.com/gokells/caflow.git
cd caflow/backend
```

2. Install dependencies:
```bash
go mod download
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your database and API keys
```

4. Run database migrations:
```bash
go run main.go migrate
```

5. Start the server:
```bash
go run main.go serve
```

### Frontend Setup

1. Navigate to the frontend directory:
```bash
cd ../frontend
```

2. Install dependencies:
```bash
npm install
```

3. Set up environment variables:
```bash
cp .env.example .env.local
# Edit .env.local with your API URL
```

4. Start the development server:
```bash
npm run dev
```

## 🐳 Docker Deployment

We provide a complete Docker setup for easy deployment:

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## 📁 Project Structure

```
cafe-management-system/
├── backend/
│   ├── cmd/
│   ├── internal/
│   │   ├── api/          # HTTP handlers
│   │   ├── service/      # Business logic
│   │   ├── repository/   # Data access layer
│   │   └── models/       # Data models
│   ├── pkg/
│   └── migrations/       # Database migrations
├── frontend/
│   ├── components/       # Reusable UI components
│   ├── pages/           # Next.js pages
│   ├── hooks/           # Custom React hooks
│   ├── services/        # API calls
│   └── styles/          # CSS styles
└── docker-compose.yml
```

## 🔌 API Documentation

Once the server is running, API documentation is available at:
- Swagger UI: `http://localhost:8080/swagger/index.html`
- Postman Collection: See `/docs/postman` directory

### Example API Calls

**Create Order:**
```bash
curl -X POST http://localhost:8080/api/orders \
  -H "Content-Type: application/json" \
  -d '{
    "type": "dine-in",
    "items": [
      {
        "menu_item_id": 1,
        "quantity": 2,
        "modifiers": ["skim-milk", "extra-shot"]
      }
    ]
  }'
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guidelines](CONTRIBUTING.md) for details.

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support, please open an issue in the GitHub repository or contact our team at support@cafemanagement.com.

## 🙏 Acknowledgments

- Toast POS for inspiration on user experience
- Square for payment processing integration patterns
- The Go and React communities for excellent tools and libraries

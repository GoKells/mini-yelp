# Project Nomad - Frontend

This repository contains the frontend application for **Project Nomad**, a modern, scalable online reviews platform similar to Yelp. It is built to provide a fast, responsive, and engaging user experience for discovering and reviewing businesses.

## 🚀 Tech Stack

- **Framework:** [React 18](https://reactjs.org/) with Functional Components & Hooks
- **Build Tool:** [Vite](https://vitejs.dev/) (for fast development and builds)
- **Language:** [TypeScript](https://www.typescriptlang.org/)
- **Styling:** [Tailwind CSS](https://tailwindcss.com/)
- **HTTP Client:** [Axios](https://axios-http.com/)
- **State Management:** [TanStack Query (React Query)](https://tanstack.com/query/latest) for server state & [Zustand](https://github.com/pmndrs/zustand) for app state
- **Routing:** [React Router v6](https://reactrouter.com/)
- **Form Handling:** [React Hook Form](https://react-hook-form.com/)
- **Testing:** [Vitest](https://vitest.dev/) + [React Testing Library](https://testing-library.com/docs/react-testing-library/intro/)
- **Linting & Formatting:** [ESLint](https://eslint.org/) + [Prettier](https://prettier.io/)

## 📁 Project Structure

```
src/
├── components/          # Reusable UI components (Button, Card, Modal)
│   ├── ui/             # Basic unstyled components
│   ├── business/       # Business-specific components (BusinessCard, ReviewList)
│   └── common/         # Other shared components (Header, Footer, SearchBar)
├── pages/              # Top-level page components (Home, Business, Profile)
├── hooks/              # Custom React hooks (useAuth, useSearch)
├── services/           # API service layer (api-client.js, businessService.js)
├── stores/             # Zustand state stores (authStore.ts, uiStore.ts)
├── types/              # TypeScript type definitions
├── utils/              # Helper functions and utilities
├── assets/             # Static assets (images, icons)
└── __tests__/          # Directory for test files
```

## 🛠️ Prerequisites

- **Node.js** (v18 or higher)
- **npm** or **yarn** or **pnpm**
- The [Project Nomad Backend](https://github.com/your-org/project-nomad-backend) must be running locally or accessible via its API URL.

## ⚡ Getting Started

1.  **Clone the repository and navigate to the frontend folder:**
    ```bash
    git clone https://github.com/your-org/project-nomad.git
    cd project-nomad/frontend
    ```

2.  **Install dependencies:**
    ```bash
    npm install
    # or
    yarn
    # or
    pnpm install
    # or
    bun install
    ```

3.  **Environment Setup:**
    Copy the example environment file and update the variables with your local configuration.
    ```bash
    cp .env.example .env.local
    ```
    Edit `.env.local`:
    ```env
    VITE_APP_NAME="Project Nomad"
    VITE_API_BASE_URL=http://localhost:3001/api/v1
    VITE_ES_PROXY_URL=http://localhost:3001/api/v1/search/proxy # Optional, for direct search dev
    ```

4.  **Start the development server:**
    ```bash
    npm run dev
    # or
    yarn dev
    # or
    pnpm dev
    # or
    bun dev
    ```
    The app will open on `http://localhost:5173`.

## 🤝 Connecting to the Backend

This frontend is designed to work with the backend API described in the main PRD. Key endpoints include:

- `GET /api/v1/businesses` - List & search businesses (via Elasticsearch)
- `GET /api/v1/businesses/:id` - Get a single business and its reviews
- `POST /api/v1/reviews` - Create a new review
- `POST /api/v1/upload` - Upload an image to S3

Ensure your backend is running on the port specified in your `VITE_API_BASE_URL` (e.g., `http://localhost:3001`).

## 📜 Available Scripts

- `npm run dev` - Starts the development server with HMR.
- `npm run build` - Builds the app for production to the `dist` folder.
- `npm run preview` - Previews the production build locally.
- `npm run test` - Runs the test suite with Vitest.
- `npm run test:ui` - Runs the test suite in UI mode.
- `npm run lint` - Runs ESLint to catch linting errors.
- `npm run type-check` - Runs the TypeScript compiler to check for type errors.

## 🧪 Testing

This project uses Vitest and React Testing Library for unit and component testing.

- Write tests in files named `*.test.ts` or `*.test.tsx` alongside your code or in the `__tests__` directory.
- Run `npm run test` to execute the test suite.

Example test:
```jsx
import { describe, it, expect } from 'vitest';
import { render, screen } from '@testing-library/react';
import BusinessCard from '../components/business/BusinessCard';

describe('BusinessCard', () => {
  it('renders the business name and rating', () => {
    const mockBusiness = { name: 'Test Cafe', average_rating: 4.5 };
    render(<BusinessCard business={mockBusiness} />);
    expect(screen.getByText('Test Cafe')).toBeInTheDocument();
    expect(screen.getByText('4.5')).toBeInTheDocument();
  });
});
```

## 🚀 Deployment

The `build` command creates a static bundle in the `dist` directory, ready to be served by any static host (Netlify, Vercel, S3, etc.).

1.  **Build the project:**
    ```bash
    npm run build
    ```

2.  **Deploy the `dist` folder to your preferred platform.**

3.  **Ensure the following environment variables are set in your production host:**
    - `VITE_API_BASE_URL` (Points to your live backend API, e.g., `https://api.myapp.com`)

## 🤔 FAQ

**Q: Why Vite instead of Create React App?**
**A:** Vite offers significantly faster server start and hot module replacement (HMR), leading to a better development experience.

**Q: Why two state management libraries?**
**A:** TanStack Query is perfect for managing asynchronous server state (caching, syncing, updating data from the API). Zustand is a simple, unopinionated library for managing client-side UI state (e.g., is a modal open?).

**Q: How do I add a new page?**
**A:** 1. Create a new component in `/pages`. 2. Add a new route in `App.tsx`. 3. (Optional) Add a link to the new page in the navigation component.

### **Product Requirements Document: "Mini Yelp"**

**Version:** 1.0
**Date:** August 24, 2025
**Author:** Kelly Mahlangu
**Status:** Final

---

### **1. Executive Summary**

This document outlines the requirements for building the backend of a modern, scalable online reviews platform. The core product will allow users to discover businesses, read and write reviews, and upload media. Our primary technical goals are to ensure high performance under heavy read load, provide robust and relevant search functionality, and maintain data consistency across distributed systems.

### **2. Problem Statement**

Users lack a trusted, centralized platform to find authentic reviews and photos for local businesses. Existing solutions can be slow, have poor search relevance, or lack real-time updated information, leading to a frustrating user experience.

### **3. Objectives & Key Results (OKRs)**

*   **Objective:** Build a reliable and engaging foundation for user-generated business reviews.
    *   **KR1:** Achieve 99.9% uptime for core read APIs in the first quarter post-launch.
    *   **KR2:** Achieve a sub-200ms average response time for business profile and review pages.
    *   **KR3:** Achieve a user satisfaction score (CSAT) of >4.5/5 for the review submission process.
*   **Objective:** Deliver a best-in-class search and discovery experience.
    *   **KR1:** Deploy a functional full-text search with faceted filtering by launch.
    *   **KR2:** 80% of users who use search click on a result from the first page.
*   **Objective:** Ensure system scalability to support rapid growth.
    *   **KR1:** Design a system capable of handling 100:1 read-to-write ratio.
    *   **KR2:** Support a database of 100k businesses and 1M+ reviews.

### **4. User Personas**

*   **Paul the Patron:** Wants to quickly find a great restaurant. Needs fast, relevant search results, easy-to-skim reviews, and high-quality photos.
*   **Rachel the Reviewer:** Enjoys sharing her experiences. Needs a seamless, fast process for writing reviews and uploading photos. Values having her activity acknowledged.
*   **Business Owner Ben:** Monitors his business's reputation. Needs to see new reviews quickly and get accurate average rating data.

### **5. Core Features & User Stories**

#### **Epic: Business Discovery & Reviews**

*   **User Story:** As Paul the Patron, I want to search for businesses by name, cuisine, and location so that I can find relevant options quickly.
    *   **Acceptance Criteria:**
        *   AC1: Search query returns results ranked by relevance in <500ms.
        *   AC2: Results can be filtered by average rating (e.g., 4 stars and up).
        *   AC3: Results display the business name, average rating, and number of reviews.
*   **User Story:** As Paul the Patron, I want to view a business profile with all its reviews and average rating so I can decide if I want to go there.
    *   **Acceptance Criteria:**
        *   AC1: Business profile page loads in <200ms.
        *   AC2: Average rating is displayed and is mathematically correct.
        *   AC3: Reviews are paginated (20 per page).
        *   AC4: Associated photos are displayed in a gallery.
*   **User Story:** As Rachel the Reviewer, I want to write a review with a star rating and text and upload photos so I can share my experience.
    *   **Acceptance Criteria:**
        *   AC1: A user can only submit one review per business.
        *   AC2: The UI allows uploading up to 5 photos per review.
        *   AC3: The system confirms the review was submitted successfully.
        *   AC4: The business's average rating updates within 5 seconds of the review being posted.
        *   AC5: The review becomes immediately visible on the business page.

#### **Epic: Media Management**

*   **User Story:** As the System, I want to securely store user-uploaded images and serve them efficiently.
    *   **Acceptance Criteria:**
        *   AC1: All images are stored in S3.
        *   AC2: Only authenticated users can upload images.
        *   AC3: Image links stored in the database are immutable and point to the correct S3 resource.

### **6. Technical Architecture & Implementation**

**6.1. High-Level System Design**
The system will use a decoupled, event-driven architecture to ensure scalability and reliability for its read-heavy workload.

*   **Primary Datastore:** PostgreSQL RDS instance to serve as the source of truth for users, businesses, and reviews.
*   **Search Engine:** Elasticsearch (ES) cluster to handle all search queries, providing fast and relevant full-text search.
*   **Object Storage:** Amazon S3 for storing all user-uploaded images. Links to these images will be stored in PostgreSQL.
*   **Message Queue:** RabbitMQ or Amazon SQS to decouple the primary application from the search index更新 process.
*   **Application Layer:** Dockerized Node.js API servers to handle HTTP requests and business logic.

**6.2. Data Flow & Key Processes**

*   **Write Path (Review Submission):**
    1.  User submits a review via the API.
    2.  API validates input, stores the review in PostgreSQL, and updates the precomputed `average_rating` and `review_count` on the `businesses` table.
    3.  API uploads images to S3 and stores the links in the `review_images` table.
    4.  API publishes an event to the message queue containing the review data (e.g., `review_id`, `business_id`, text).
    5.  **→ Success response is sent to the user.**
    6.  A separate **Indexer Service** consumes events from the queue and updates the corresponding document in the Elasticsearch cluster.

*   **Read Path (Search):**
    1.  User enters a search query.
    2.  API forwards the query to the Elasticsearch cluster.
    3.  ES returns sorted, relevant results.
    4.  API returns results to the client.

*   **Read Path (Business Page):**
    1.  User requests a business profile.
    2.  API fetches business data (including precomputed averages) from PostgreSQL.
    3.  API fetches paginated reviews and photo links from PostgreSQL.
    4.  API returns the compiled data to the client.

### **7. Non-Functional Requirements**

*   **Performance:** 95% of API responses must be under 200ms. Search queries must respond in <500ms.
*   **Availability:** The core service (reading reviews/businesses) must have 99.9% uptime.
*   **Scalability:** The system must be designed to scale horizontally for reads (API layer, ES cluster) and handle a write throughput of 100 reviews/second.
*   **Security:** All endpoints must be authenticated where appropriate (e.g., writing reviews). Data in transit (TLS) and at rest (encryption) must be encrypted. S3 buckets must have strict, non-public access policies.

### **8. Dependencies & Constraints**

*   **Dependencies:** AWS infrastructure (S3, RDS, ECS/EKS if used), Elasticsearch cluster management, CI/CD pipeline.
*   **Constraints:** Initial team size of 5 engineers. MVP must be launched within 6 months.
*   **Risks & Mitigation:**
    *   **Risk:** Eventual consistency between PostgreSQL and Elasticsearch could cause temporary search discrepancies.
    *   **Mitigation:** Implement a dead-letter queue for failed index updates and a dashboard to monitor queue length and processing lag. Clearly communicate this behavior in the UI if necessary (e.g., "new reviews may take a moment to appear in search").
    *   **Risk:** Precomputing averages on write could become a bottleneck if a single business receives a huge volume of reviews simultaneously.
    *   **Mitigation:** Use database transactions and optimistic concurrency control to handle concurrent updates to the `businesses` table.

### **9. KPIs & Metrics**

*   **Business Metrics:** Number of active users, number of reviews per day, number of businesses added per week.
*   **Performance Metrics:** P95 API latency, Elasticsearch query latency, message queue backlog.
*   **Reliability Metrics:** Uptime %, error rate (5xx responses).

### **10. Open Questions & Future Considerations**

*   How will we handle review moderation for spam and inappropriate content?
*   How will we manage database schema migrations?
*   **Future Features:** User social graphs (followers), reactions to reviews, ad serving platform, real-time notifications, advanced analytics for business owners.

---

### **Analysis of Improvements & Rationale**

1.  **Clarity & Structure:** The original script was a brain dump. I organized it into a standard PRD format with clear sections (Executive Summary, OKRs, User Stories, Tech Specs). This allows engineers, designers, and executives to quickly find the information relevant to them.

2.  **Completeness:**
    *   Added **OKRs** to align the team on measurable goals.
    *   Defined **User Personas** to ground feature discussions in user needs.
    *   Broke down features into detailed **User Stories** with specific **Acceptance Criteria** (e.g., pagination, concurrency, upload limits). This eliminates ambiguity for developers.
    *   Added a dedicated **Technical Architecture** section that details the data flow for both writes and reads, which was only vaguely described.
    *   Defined **Non-Functional Requirements** (Performance, Security) which are critical for a high-quality product.
    *   Explicitly listed **Risks and Mitigations**, which is crucial for technical planning.

3.  **Impact & Value:**
    *   The focus on performance (sub-200ms response) directly impacts user satisfaction and engagement.
    *   The event-driven architecture with a queue ensures the core application remains responsive even if the search indexer is slow, improving overall system resilience and user experience.

4.  **Feasibility & Technical Insight:**
    *   **Precomputing Averages:** Agreed with the original assessment. This is the right choice for a read-heavy system. I added a risk about concurrent writes to the same business, which is a realistic edge case (e.g., a viral business).
    *   **S3 for Media:** This is a best practice. I strengthened it by adding security considerations (non-public buckets).
    *   **Elasticsearch for Search:** Correctly identified the need. I formalized the process by introducing a message queue, which is a more robust and scalable pattern than having the API update ES directly. This prevents the main API from being slowed down by ES performance issues.
    *   **Specificity:** Changed "Postgress" to "PostgreSQL RDS" and suggested specific queue technologies (RabbitMQ/SQS), moving from vague concepts to actionable implementation details.

5.  **Language & Precision:**
    *   Changed _"perform some logic"_ to a detailed, step-by-step data flow.
    *   Changed _"store it in Postgress"_ to _"store the immutable image link in the `review_images` table in PostgreSQL."_
    *   Made requirements **measurable** (e.g., "<200ms", "99.9% uptime").

6.  **Prioritization:**
    *   **Must-Have (MVP):** Core tables (Users, Businesses, Reviews), Review CRUD API, Precomputed averages, Image upload to S3, Elasticsearch integration with queue for search.
    *   **Should-Have:** Pagination, advanced filters (rating, cuisine), comprehensive monitoring and metrics dashboards.
    *   **Nice-to-Have:** Review moderation tools, real-time notifications, user profile pages, analytics for business owners.

7.  **Optional Innovations:**
    *   **Photo Moderation Service:** A future AI-based service to blur license plates or detect inappropriate imagery automatically would build trust and safety.
    *   **Offline-First Review Drafting:** Allowing users to start a review without a network connection and sync it later could significantly increase review completion rates.
    *   **"Review Summary" NLP:** Using Natural Language Processing on reviews to generate a summary of common pros and cons (e.g., "Great for groups," "Noisy atmosphere") would provide immense value to users skimming reviews.

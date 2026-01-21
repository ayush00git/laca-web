# LACA Registration Portal - NIT Hamirpur

The official website for **LA/CA (Lecture Added Courses/Co-Curricular Activities)** registration at **NIT Hamirpur**. This portal enables students to register for elective activities and courses with a seamless, scalable experience.

## 📊 Performance Stats

- **850+** registrations completed in **2 days**
- **60,000+** HTTP requests handled
- **99.9%** uptime during peak registration period
- Efficient email delivery with Redis-based job queuing

Read the full technical blog post: [Blog](https://tnc.ayushz.me/blogs/695eac9fb6ba0a3261cea389)

---

## 🛠️ Tech Stack

### Backend
- **Node.js** - JavaScript runtime
- **Express.js** - Web framework for routing and middleware
- **MongoDB** - NoSQL database for storing registrations and activity data
- **Mongoose** - ODM (Object Data Modeling) for MongoDB

### Email & Job Queue
- **BullMQ** - Task queue built on Redis for background job processing
- **IORedis** - Redis client for Node.js
- **Nodemailer** - Email sending with SMTP support (Gmail)
- **Redis** - In-memory data store for job queue and worker coordination

### Frontend
- **EJS** - Embedded JavaScript templating
- **HTML5 & CSS3** - Modern, responsive UI
- **Vanilla JavaScript** - Form handling and dynamic subject loading

### Authentication & Security
- **JWT (jsonwebtoken)** - Token-based email verification
- **dotenv** - Environment variable management

---

## 📋 Project Structure

```
laca-site/
├── models/
│   ├── form.js              # User registration schema
│   └── subject.js           # Activity/course schema
├── routes/
│   ├── form.js              # Form submission & validation
│   ├── email.js             # Email verification endpoint
│   ├── subject.js           # Subject management & API
│   └── details.js           # Registration details (admin)
├── services/
│   ├── sendEmail.js         # Email template & transport logic
│   └── redis/
│       ├── emailQueue.js    # BullMQ queue producer
│       └── emailWorker.js   # BullMQ worker consumer
├── views/
│   ├── home.ejs             # Main registration form
│   ├── verify-email.ejs     # Email verification page
│   ├── subject.ejs          # Admin subject management
│   └── details.ejs          # Registration details page
├── public/                  # Static assets
├── index.js                 # Server entry point
├── connection.js            # MongoDB connection handler
├── package.json             # Dependencies
└── README.md
```

---

## 🚀 Getting Started (Local Setup)

### Prerequisites
- **Node.js** (v14 or higher)
- **npm** (comes with Node.js)
- **MongoDB** (local or cloud instance via MongoDB Atlas)
- **Redis** (for background job processing)
- **Git**

### 1. Clone the Repository

```bash
git clone https://github.com/ayush00git/laca-site.git
cd laca-site
```

### 2. Install Dependencies

```bash
npm install
```

### 3. Configure Environment Variables

Create a `.env` file in the project root:

```env
# Server
PORT=8089
NODE_ENV=development

# Database
MONGO_URI=mongodb://localhost:27017/laca-site
# Or use MongoDB Atlas:
# MONGO_URI=mongodb+srv://username:password@cluster.mongodb.net/laca-site?retryWrites=true&w=majority

# Email Configuration (Gmail SMTP)
USER_MAIL=your-email@gmail.com
USER_PASS=your-app-specific-password
# Generate app password: https://myaccount.google.com/apppasswords

# JWT & Verification
JWT_SECRET=your-super-secret-key
PROD_URL=http://localhost:8089

# Redis Configuration (for email queue)
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
# Or use Redis URL:
# REDIS_URL=redis://:password@host:port

# Private Routes (admin/sensitive operations)
PRIVATE_URL=admin
```

### 4. Start MongoDB

#### Option A: Local MongoDB
```bash
# If installed locally
mongod
```

#### Option B: MongoDB Atlas (Cloud)
- Create a cluster at [MongoDB Atlas](https://www.mongodb.com/cloud/atlas)
- Get connection string and add to `.env` as `MONGO_URI`

### 5. Start Redis Server

```bash
# If installed locally
redis-server

# Or using Docker
docker run -d -p 6379:6379 redis:latest
```

### 6. Start the Background Worker (Email Processing)

In a **separate terminal**:

```bash
npm run start-worker
```

This will start the BullMQ worker that processes email jobs from the Redis queue.

### 7. Start the Development Server

In another terminal:

```bash
npm run dev
```

The server will start at **http://localhost:8089**

---

## 📧 Email Queue & Redis Architecture

### How Email Delivery Works

1. **User Submits Form** → Server validates & creates JWT token
2. **Enqueue Job** → Job is added to Redis queue (`emailQueue.js`)
3. **Immediate Response** → User sees success message instantly (non-blocking)
4. **Worker Processes** → Background worker (`emailWorker.js`) picks up job from queue
5. **Email Sent** → Nodemailer sends verification email via Gmail SMTP
6. **Job Complete** → Job is marked complete and removed from queue (or retained on failure)

## 📝 API Endpoints

### Registration
- `POST /form/submit` - Submit registration form
- `GET /email/verify-email?token=JWT_TOKEN` - Verify email and confirm registration

### Subject Management
- `GET /admin/subject/api/all` - Fetch all activities (for frontend dropdown)
- `POST /admin/subject` - Create new activity (admin only)

### Admin
- `GET /admin/details` - View all registrations

### Config
- `GET /api/config` - Get dynamic API configuration

---

## 🔄 Available npm Scripts

```bash
npm run dev              # Start development server with auto-reload (nodemon)
npm run start            # Start production server
npm run start-worker     # Start background email worker
```

---

## 🛡️ Security Features

- **JWT-based email verification** - Tokens expire in 24 hours
- **Input validation** - Server-side validation for all form inputs
- **Email domain restriction** - Only `@nith.ac.in` emails accepted
- **Rate limiting** - BullMQ limiter (10 emails per second max)
- **Environment variables** - Sensitive data not in code
- **HTTPS ready** - Configure in production deployment

---
### Environment Configuration (Production)

```env
NODE_ENV=production
PORT=8089
MONGO_URI=mongodb+srv://...           # Atlas connection
REDIS_URL=redis://:password@host:port # Cloud Redis
USER_MAIL=registration@nith.ac.in     # Official email
JWT_SECRET=long-random-secret-key
PROD_URL=https://laca.nith.ac.in      # Production domain
PRIVATE_URL=admin
```

### Recommended Hosting
- **Server:** AWS EC2.
- **Database:** MongoDB Atlas
- **Redis:** AWS ElastiCache or Upstash
- **Email:** Gmail SMTP or SendGrid/AWS SES for higher limits

### Process Manager (Production)

Use **PM2** to keep the server and worker running:

```bash
npm install -g pm2

# Start both processes
pm2 start index.js --name "laca-server"
pm2 start services/redis/emailWorker.js --name "email-worker"

# Monitor
pm2 monit

# Logs
pm2 logs
```

---

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit changes (`git commit -m 'Add amazing feature'`)
4. Push to branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

---

## 📞 Support & Contact

**Developers:**
- [Ayush](https://github.com/ayush00git)
- [Aryan](https://github.com/pixel-Neko)


## 📄 License

This project is part of NIT Hamirpur's official infrastructure. Contact the administration for licensing details.

---
**Last Updated:** January 11, 2026  
**Version:** 2.0.0

// Submitting to the college //
// DONE //

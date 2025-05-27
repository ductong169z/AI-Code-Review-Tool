# Guide to Build and Run Gemini AI Code Review CLI

This guide explains how to build and run the Go-based Gemini AI code review tool.

---

## Prerequisites

- **Go** 1.20 or higher installed: [https://golang.org/dl/](https://golang.org/dl/)
- **Git** installed and configured: [https://git-scm.com/downloads](https://git-scm.com/downloads)
- Valid **Google Gemini API Key**

---

## Setup

1. **Clone or download** the repository containing `main.go`, `.env`, and `prompt.txt`.

2. **Create `.env` file** in the same directory as `main.go` with the following content:

   ```env
   GEMINI_API_KEY=your_google_gemini_api_key_here
   ```

3. **Customize `prompt.txt`** if needed to adjust AI instructions.

4. **Build the tool**:

```bash
go build -buildvcs=false -o review-tool/reviewer main.go
```


5. **Run the tool**:
Copy the `review-tool` directory to your project root.

```bash
cd review-tool
./review.sh
```


   
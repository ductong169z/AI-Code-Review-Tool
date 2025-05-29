
# 🤖 AI Code Review Tool

An automated code review tool using AI (Google Gemini or OpenAI GPT). The `reviewer` binary helps you review your staged Git changes before pushing.

---

## 📁 Directory Structure

```
review-tool/
├── reviewer            # Review binary
├── prompt.txt        # Custom prompt for the AI reviewer
├── .env              # Contains GEMINI_API_KEY or OPENAI_API_KEY
├── review-<timestamp>.md # Output file with AI review results
```

---

## 🛠 Requirements

- Git
- API Key:
  - Google Gemini (`GEMINI_API_KEY`)
  - or OpenAI (`OPENAI_API_KEY`)
- `.env` file in the `review-tool` folder

---

## 📝 Create `.env`

In the `review-tool` folder, create a `.env` file:

```
GEMINI_API_KEY=your_google_gemini_api_key
```

Or for OpenAI:

```
OPENAI_API_KEY=your_openai_api_key
```

---

## 🚀 How to Use

You **must** change into the `review-tool` directory first:

```bash
cd review-tool
./reviewer
```

You can optimize prompt.txt to make the review more accurate.

---

## ✅ Notes

- Only works on **staged changes** (`git add`).
- If there are no changes, it will show: "No changes to review".
- Each run write `review-<timestamp>.md`.

---

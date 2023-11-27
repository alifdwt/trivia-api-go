# CRUD TRIVIA GAME KEREN GO ECHO

## Installation

1. Install Go, Postgres, Echo
2. Clone this repository

```
git clone https://github.com/alifdwt/trivia-game-keren
cd trivia-game-keren-go-echo
```

3. Configure .env file
4. Run `go run main.go`

## Documentation

### API Endpoints

#### Question

##### Get All Questions

- URL: `/api/question`
- Method: GET
- Description: Get all questions
- Require token: `No`
- Response

```
{
    "code": 200,
    "data": [
        {
            "id": 1,
            "question": "Apa nama sungai terpanjang di Jawa Barat",
            "answer": "Sungai Citarum",
            "wrong_answer_1": "Sungai Cisadane",
            "wrong_answer_2": "Sungai Ciliwung",
            "wrong_answer_3": "Sungai Cigenang",
            "created_at": "2023-11-26 20:47:06",
            "updated_at": "2023-11-26 20:47:06"
        },
        {
            "id": 2,
            "question": "Di provinsi manakah Candi Borobudur terletak?",
            "answer": "Jawa Tengah",
            "wrong_answer_1": "Yogyakarta",
            "wrong_answer_2": "Jawa Timur",
            "wrong_answer_3": "Bali",
            "created_at": "2023-11-27 10:41:07",
            "updated_at": "2023-11-27 10:42:33"
        }
    ]
}
```

##### Get Question by ID

- URL: `/api/question/:questionId`
- Method: GET
- Description: Get question by ID
- Require token: `No`
- Response

```
{
    "code": 200,
    "data": {
        "id": 1,
        "question": "Apa nama sungai terpanjang di Jawa Barat",
        "answer": "Sungai Citarum",
        "wrong_answer_1": "Sungai Cisadane",
        "wrong_answer_2": "Sungai Ciliwung",
        "wrong_answer_3": "Sungai Cigenang",
        "created_at": "2023-11-26 20:47:06",
        "updated_at": "2023-11-26 20:47:06"
    }
}
```

##### Create Question

- URL: `/api/question`
- Method: POST
- Description: Create question
- Require token: `Yes`
- Request

```
{
    "question": "Apa nama selat yang memisahkan Rusia dan Amerika Serikat?",
    "answer": "Selat Bering",
    "wrong_answer_1": "Selat Bosphorus",
    "wrong_answer_2": "Selat Malaka",
    "wrong_answer_3": "Selat Gibraltar"
}
```

- Response

```
{
    "code": 200,
    "data": {
        "id": 3,
        "question": "Apa nama selat yang memisahkan Rusia dan Amerika Serikat?",
        "answer": "Selat Bering",
        "wrong_answer_1": "Selat Bosphorus",
        "wrong_answer_2": "Selat Malaka",
        "wrong_answer_3": "Selat Gibraltar",
        "created_at": "2023-11-27 10:41:07",
        "updated_at": "2023-11-27 10:41:07"
    }
}
```

##### Update Question

- URL: `/api/question/:questionId`
- Method: PUT
- Description: Update question
- Require token: `Yes`
- Request

```
{
    "question": "Apa nama selat yang memisahkan Rusia dan Amerika Serikat?",
    "answer": "Selat Bering",
    "wrong_answer_1": "Selat Bosphorus",
    "wrong_answer_2": "Selat Sunda",
    "wrong_answer_3": "Selat Gibraltar"
}
```

- Response

```
{
    "code": 200,
    "data": {
        "id": 3,
        "question": "Apa nama selat yang memisahkan Rusia dan Amerika Serikat?",
        "answer": "Selat Bering",
        "wrong_answer_1": "Selat Bosphorus",
        "wrong_answer_2": "Selat Sunda",
        "wrong_answer_3": "Selat Gibraltar",
        "created_at": "2023-11-27 10:41:07",
        "updated_at": "2023-11-27 10:42:33"
    }
}
```

##### Delete Question

- URL: `/api/question/:questionId`
- Method: DELETE
- Description: Delete question
- Require token: `Yes`
- Response

```
{
    "code": 200,
    "data": {
        "id": 3,
        "question": "Apa nama selat yang memisahkan Rusia dan Amerika Serikat?",
        "answer": "Selat Bering",
        "wrong_answer_1": "Selat Bosphorus",
        "wrong_answer_2": "Selat Sunda",
        "wrong_answer_3": "Selat Gibraltar",
        "created_at": "2023-11-27 10:41:07",
        "updated_at": "2023-11-27 10:42:33"
    }
}
```

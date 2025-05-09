# distributed_calculator
Распределённый вычислитель арифметических выражений

#Как его запустить?
>1 Скачиваете или кланируете мой гит
>2 Пишете go run .cmd/calc_service.main.go
>P.S если у вас не получается запустить из-за CGO_ENABLED=0 хотя у вас установлен MINGW то попробуйте запустить run.bat

#Как им пользоваться?
>**1 Регистрация**
>>1.1 для начала открываем повершелл и вводим логин и пароль в accauntInfo, а именно ``` $accountInfo = @{
     login    = "yourlogin"
     password = "yourpassword"
 } | ConvertTo-Json```
>>
>>1.2 вводим команду на регистрацию ```Invoke-RestMethod -Uri "http://localhost:8080/api/v1/register" `
                   -Method POST `
                   -Body $accountInfo `
                   -ContentType "application/json"```
>
>**2 Вход**
>>
>>2.1 Так как у нас уже есть тело ```$accountInfo``` то мы вводим запрос на вход ```Invoke-RestMethod -Uri "http://localhost:8080/api/v1/login" `
                   -Method POST `
                   -Body $accountInfo `
                   -ContentType "application/json"```
>>
>>2.2 Мы должны получить ответ похожий на ```token ----- yourtoken # сам токен```
>
>
>**3 Запрос на решение примера**
>>
>>
>>
>>3.1 Сначало нам нужно ввести тело ```$headers``` где будет наш токен авторизации. Вводим ```$headers = @{
    Authorization = "Bearer yourtoken"  # вставь сюда свой полный токен
}``` Самое важное чтобы в кавычках перед токеном стоял 'Bearer'
>>
>>
>>3.2Дальше вводим запрос, ```Invoke-RestMethod -Uri "http://localhost:8080/api/v1/calculate" `
                   -Method POST `
                   -Headers $headers `
                   -Body '{"expression": "2 + 2"}' `
                   -ContentType "application/json"```
>>
>>3.3И мы должны будем получить ответ похожий на ```result ----- 4```
>>
>Теперь мы получили наш **долгожданный ответ**!

Я очень старался и надеюсь что вы поставите хороший балл!

По вопросам писать в тг **```Olomadness```**

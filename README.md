# GeoCoder
Микросервис, позволяющий получить полный адрес по координатам или неполному адресу при помощи API Яндекс Карт. Данный микросервис может использоваться в качестве автодополнения при поиске адреса, заполнения анкет. 

## Установка
* Указать операционную систему и архитектуру в Makefile
* Выполнить в папке проекта команду `make build`, получить app файл
* Создать файл `.env` в формате `example.env`: указать свободный порт, ключ API Яндекс Карт. Поместить `.env` рядом с app.
* Запустить ./app 

## Протестировать микросервис можно по адресу http://gogeocoder.ru
Пример GET запроса к микросервису для определения полного адреса по координатам:

```
curl --request GET \
  --url 'http://gogeocoder.ru:4444/reverse-geocode?lat=55.7361783299725&lng=37.62344576776365' \
  --header 'Authorization: caaba14c-c123-460e-b48b-f3d77749026a'
```
 
Ответ: 
```
{"address":"Russian Federation, Moscow, Bolshaya Ordynka Street, 42",
 "detailedAddress":{"FormattedAddress":"Russian Federation, Moscow, Bolshaya Ordynka Street, 42",
                     "Street":"Bolshaya Ordynka Street",
                     "HouseNumber":"42",
                     "Suburb":"",
                     "Postcode":"119017",
                     "State":"Moscow",
                     "StateCode":"",
                     "StateDistrict":"",
                     "County":"",
                     "Country":"Russian Federation",
                     "CountryCode":"RU",
                     "City":"Moscow"},
 "error":null,
 "message":"success",
 "status":true}
```

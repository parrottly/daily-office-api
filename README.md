# Daily Office API
API for The Book of Common Prayer's Daily Office Lectionary. I used [@reubenlillie's](https://github.com/reubenlillie/daily-office) JSON files for this project. 

[https://www.dolapi.com](https://www.dolapi.com/)

Example: [First Day of Advent](https://www.dolapi.com/year-one/advent/1/Sunday)

### Year
Retrieves all readings for a the given year. Holy days and special occasions are also included in this.
- year-one
- year-two
- holy-days
- special occasions

ex: `GET https://www.dolapi.com/year-two`

### Season
- advent
- christmas
- epiphany
- lent 
- easter
- after-pentecost

ex: `GET https://www.dolapi.com/year-two/after-pentecost`


### Week
Retrieves all readings for a given week of a season. Represented as a number.

ex: `GET https://www.dolapi.com/year-two/after-pentecost/3`


### Day
All readings for a given day of the week.

ex: `GET https://www.dolapi.com/year-two/after-pentecost/3/thursday`

### Psalms
Only retrieve Psalms for a given day.

ex: `GET https://www.dolapi.com/year-two/after-pentecost/3/thursday/psalms`


### Lessons
Only retrieve lessons for a given day.

ex: `GET https://www.dolapi.com/year-two/after-pentecost/3/thursday/lessons`


I mainly built this API to learn Go but if any of you Episcopalian devs out there (there are dozens of us!) find this useful and have feedback, feel free to open an issue or create a PR. May the peace of the Lord be always with you!


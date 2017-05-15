//TODO: implement the zip server in Node.js!

"use strict";

const express = require('express')
const cors = require('cors')
const morgan = require('morgan')
const zips = require('../data/zips.json')
console.log('loaded %d zips', zips.length)
console.log(`loaded ${zips.length}`)
// const mymod = require("./mymodule.js")
// const data = require("./mydata.json")

// let zipCityIndex;

// for zip, z := range {
// cityLower = zips.cityLower() 
//     zipCityIndex.append(zip, z)
// }

const zipCityIndex = zips.reduce((index, record) => {
    let cityLower = record.city.toLowerCase();
    let zipsForCity = index[cityLower]
    if (!zipsForCity) {
        index[cityLower] = zipsForCity = [];
    }
    zipsForCity.push(record)
    return index;
}, {})
console.log(`there are ${zipCityIndex.seattle.length}`)

const app = express();

const port = process.env.PORT || 80;
const host = process.env.HOST || '';

app.use(morgan('dev'))
app.use(cors())

app.get('/zips/city/:cityName', (req, res) => {
    let zipsForCity = zipCityIndex[req.params.cityName.toLowerCase()]
    if (!zipsForCity) {
        res.status(404).send('invalid city name')
    } else {
        res.json(zipsForCity)
    }
})


app.get('/hello/:name', (req, res) => {
    res.send(`Hello ${req.params.name}!`)
})

app.listen(port, host, () => {
    console.log(`server is listening at http://${host}:${port}`)
})
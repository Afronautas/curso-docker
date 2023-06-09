"use strict"

const express = require("express")

const PORT = 3000
const HOST = "0.0.0.0"

const app = express()

app.get("/", (req, res) => {
    res.send("Hello, Docker!")
})

app.listen(PORT, HOST, () => {
    console.log(`Running server on http://${HOST}:${PORT}`)
})
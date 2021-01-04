import axios from "axios";
import express from "express"
import mongoose from "mongoose"

const app = express()

const {
    MONGO_URL,
    MONGO_DATABASE,
    MONGO_USERNAME,
    MONGO_PASSWORD,
    AUTH_URL,
    SECURE_URL
} = process.env;

app.get('/', async (req, res, next) => {
    res.status(200).send('Test Microservice is up!')
})

app.get('/health', async (req, res, next) => {
    var resp = {
        secureHealthy: false,
        authHealthy: false
    }
    axios.get(`${AUTH_URL}`).then(() => { resp.authHealthy = true }).catch()
    axios.get(`${SECURE_URL}`).then(() => { resp.secureHealthy = true }).catch()
    res.status(200).json(resp)
})

mongoose.connect(
    `mongodb://${MONGO_URL}/${MONGO_DATABASE}`, {
    useNewUrlParser: true,
    useUnifiedTopology: true,
    auth: {
        user: MONGO_USERNAME as string,
        password: MONGO_PASSWORD as string
    }
}
).then(() => {
    app.listen(5000, () => {
        console.log('Listening on port 5000')
    })
}).catch((reason) => {
    console.log(`${reason}\nApp Crashed`)
})
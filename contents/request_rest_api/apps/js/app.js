const express = require('express');
const app = express();
const port = 3004;

// Health Check API
app.get('/health/', (req, res) => {
    res.json({status: 'ok'});
});

// 서버 시작
app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}/`);
});

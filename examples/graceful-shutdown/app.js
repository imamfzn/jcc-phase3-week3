const app = require('express')();

const server = app.listen(3000, () => console.log("app is running"));

const sleep = time => new Promise(resolve => setTimeout(resolve, time));

app.get('/test', async function (req, res) {
  await sleep(15000);
  res.send('ok');
});

process.on('SIGTERM', async function () {
  server.close(function (err) {
    process.exit(0);
  });
});

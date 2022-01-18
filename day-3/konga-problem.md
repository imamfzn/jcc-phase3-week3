# Kong JWT Problem (SOLVED)

Masalah jwt yang sebelumnya adalah karena kita salah memasukan parameter `iss`. Seharusnya di simpan pada https://github.com/imamfzn/jcc-phase3-week3/blob/031b266b55114f9d3de00059b05c83e743216aa4/examples/monolithic-service-example/src/lib/jwt.js#L3-L7.

Contoh:

```js
const jwt = require('jsonwebtoken');

module.exports = {
  generate({ id, username, role }) {
    return jwt.sign({ id, username, role, iss: "auth.service" }, process.env.ACCESS_TOKEN_SECRET, { expiresIn: '1h' });
  }
};
```

Silahkan dicoba ya seperti kemarin. Good Luck!

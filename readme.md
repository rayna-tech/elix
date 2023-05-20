![Elix Banner](https://cdn.rayna.tech/elixdb-banner.png)

*A simple yet effective go lang based key value store database which has socket (WIP) and web api support.*

<br/>

### ⚠️ Notice

Currently elix is a WIP and is  guarenteed to change. It is in a stable release currently and will work yet any changes may break current apps that may use elix.
<br/>
<br/>

---
### 🐳 Self Hosting ( Docker )

Currently elix only supports web api requests ([sockets will be coming later](#plans)) which is why only one port is being exposed.

```sh
git clone https://github.com/rayna-tech/elix
cd elix
docker build -t elix .
docker run -it -p 8080:8080 --name="elix" elix
docker logs -f <PROVIDED_DOCKER_ID>
```
---
### 📑 Plans

- Sockets
    - Subscribe to all keys/singular key incase of updates or deletion etc.

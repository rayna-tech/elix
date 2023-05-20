const { ElixClient } = require("..");

const e = new ElixClient({
  url: "http://localhost:8080",
  seedData: [{
    key: "user",
    value: true
  }]
});

async function main() {
  let running = await e.get("user");
  console.log(running);
//   await e.remove("user")
//     running = await e.get("user");
//   console.log(running);
}

main();

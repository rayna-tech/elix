import { ClientParams } from "../types";
import axios from "axios";
import Logger from "../utils/Logger";

export default class Client {
  options: ClientParams;
  constructor(options: ClientParams) {
    options.url = options.url.endsWith("/")
      ? options.url
          .split("/")
          .slice(0, options.url.split("/").length - 1)
          .join("/")
      : options.url;

    this.options = options;
    if (this.options.seedData) {
      this.seed();
    }
  }

  async set(key: string, value: any) {
    const req = await axios
      .post(
        `${this.options.url}/new`,
        JSON.stringify({
          key: key,
          store: value,
        })
      )
      .then((r) => {
        return {
          ok: true,
          data: r.data,
        };
      })
      .catch((r) => {
        return {
          ok: false,
          reason: `${r}`,
        };
      });
    if (!req.ok) {
      //@ts-ignore
      Logger(
        `Error setting key value. Please check the db server terminal for more information.`
      ).error();
    } else {
      Logger(
        `Key: \"${key}\" has been assigned to the value: ${JSON.stringify(value)}`
      );
    }
  }

  async get(key: string): Promise<any> {
    const req = await axios
      .get(`${this.options.url}/get/${key}`)
      .then((r) => {
        return {
          ok: true,
          data: r.data,
        };
      })
      .catch((r) => {
        return {
          ok: false,
          reason: `${r}`,
        };
      });
    if (!req.ok) {
      //@ts-ignore
      Logger(
        `Error getting key value. Please check the db server terminal for more information.`
      ).error();
    } else {
      //@ts-ignore
      return req.data;
    }
  }

  async wipe() {
    const req = await axios
      .delete(`${this.options.url}/clear`)
      .then((r) => {
        return {
          ok: true,
        };
      })
      .catch((r) => {
        return {
          ok: false,
        };
      });
    if (!req.ok) {
      Logger(
        `Error clearing/wiping store. Please check the db server terminal for more information.`
      ).error();
    } else {
      Logger(`Elix store has been cleared/wiped.`).info();
    }
  }

  async remove(key: string) {
    const req = await axios
      .delete(`${this.options.url}/delete?key=${key}`)
      .then((r) => {
        return {
          ok: true,
          data: r.data,
        };
      })
      .catch((r) => {
        return {
          ok: false,
          reason: r,
        };
      });
    if (!req.ok) {
      Logger(
        `Error deleted pair in store. Please check the db server terminal for more information.`
      ).error();
    } else {
      Logger(`Pair: \"${key}\" has been deleted.`).info();
    }
  }

  async update(key: string, value: any) {
    const req = await axios
      .post(
        `${this.options.url}/update`,
        JSON.stringify({
          key: key,
          store: value,
        })
      )
      .then((r) => {
        return {
          ok: true,
          data: r.data,
        };
      })
      .catch((r) => {
        return {
          ok: false,
          reason: `${r}`,
        };
      });
    if (!req.ok) {
      //@ts-ignore
      Logger(
        `Error reassigning key value. Please check the db server terminal for more information.`
      ).error();
    } else {
      Logger(
        `Key: \"${key}\" has had a new value assigned to it.`
      );
    }
  }
  seed() {
    this.options.seedData.forEach((seed) => {
      this.set(seed.key, seed.value);
    });
  }
}

import { KeyValuePair } from "../types";
export type ClientParams = {
    url: string;
    seedData?: KeyValuePair[];
};
export default class Client {
    options: ClientParams;
    constructor(options: ClientParams);
    set(key: string, value: any): Promise<void>;
    get(key: string): Promise<any>;
    wipe(): Promise<void>;
    remove(key: string): Promise<void>;
    seed(): void;
}

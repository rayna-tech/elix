export type Object<V=any> = {
    [key: string]: V
}

export type KeyValuePair = {
    key: string,
    value: any
}

export type ClientParams = {
    url: string;
    seedData?: KeyValuePair[];
  };
  
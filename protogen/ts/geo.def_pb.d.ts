import * as jspb from "google-protobuf"

export class Location extends jspb.Message {
  getCountry(): string;
  setCountry(value: string): void;

  getRegion(): string;
  setRegion(value: string): void;

  getCity(): string;
  setCity(value: string): void;

  getPostalCode(): string;
  setPostalCode(value: string): void;

  getGeoLat(): number;
  setGeoLat(value: number): void;

  getGeoLng(): number;
  setGeoLng(value: number): void;

  getGeoPrecisionRad(): number;
  setGeoPrecisionRad(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Location.AsObject;
  static toObject(includeInstance: boolean, msg: Location): Location.AsObject;
  static serializeBinaryToWriter(message: Location, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Location;
  static deserializeBinaryFromReader(message: Location, reader: jspb.BinaryReader): Location;
}

export namespace Location {
  export type AsObject = {
    country: string,
    region: string,
    city: string,
    postalCode: string,
    geoLat: number,
    geoLng: number,
    geoPrecisionRad: number,
  }
}


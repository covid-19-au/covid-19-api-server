import * as jspb from "google-protobuf"

export class NamedLocation extends jspb.Message {
  getCountry(): string;
  setCountry(value: string): void;

  getRegion(): string;
  setRegion(value: string): void;

  getCity(): string;
  setCity(value: string): void;

  getPostalCode(): string;
  setPostalCode(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): NamedLocation.AsObject;
  static toObject(includeInstance: boolean, msg: NamedLocation): NamedLocation.AsObject;
  static serializeBinaryToWriter(message: NamedLocation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): NamedLocation;
  static deserializeBinaryFromReader(message: NamedLocation, reader: jspb.BinaryReader): NamedLocation;
}

export namespace NamedLocation {
  export type AsObject = {
    country: string,
    region: string,
    city: string,
    postalCode: string,
  }
}

export class GPSLocation extends jspb.Message {
  getGeoLat(): number;
  setGeoLat(value: number): void;

  getGeoLng(): number;
  setGeoLng(value: number): void;

  getPrecisionRad(): number;
  setPrecisionRad(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GPSLocation.AsObject;
  static toObject(includeInstance: boolean, msg: GPSLocation): GPSLocation.AsObject;
  static serializeBinaryToWriter(message: GPSLocation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GPSLocation;
  static deserializeBinaryFromReader(message: GPSLocation, reader: jspb.BinaryReader): GPSLocation;
}

export namespace GPSLocation {
  export type AsObject = {
    geoLat: number,
    geoLng: number,
    precisionRad: number,
  }
}


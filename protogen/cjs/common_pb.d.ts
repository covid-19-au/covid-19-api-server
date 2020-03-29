import * as jspb from "google-protobuf"

export class Pagination extends jspb.Message {
  getTotalResults(): number;
  setTotalResults(value: number): void;

  getTotalPages(): number;
  setTotalPages(value: number): void;

  getCurrentPage(): number;
  setCurrentPage(value: number): void;

  getNumPerPage(): number;
  setNumPerPage(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Pagination.AsObject;
  static toObject(includeInstance: boolean, msg: Pagination): Pagination.AsObject;
  static serializeBinaryToWriter(message: Pagination, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Pagination;
  static deserializeBinaryFromReader(message: Pagination, reader: jspb.BinaryReader): Pagination;
}

export namespace Pagination {
  export type AsObject = {
    totalResults: number,
    totalPages: number,
    currentPage: number,
    numPerPage: number,
  }
}


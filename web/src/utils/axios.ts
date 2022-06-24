import Axios from "axios";

async function decodeEthRawTransaction(raw: string) {
  return Axios.post("/api/decode-eth-raw", {
    raw_tx: raw,
  });
}

export { decodeEthRawTransaction };

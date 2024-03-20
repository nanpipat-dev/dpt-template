import {randomUUID} from "crypto";

export const getTimeUTC = () => new Date().toISOString().slice(0, -1);

export const getUUID = () => randomUUID()

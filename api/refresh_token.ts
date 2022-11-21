import type {VercelRequest, VercelResponse} from "@vercel/node";
import {AuthorizationCode} from "simple-oauth2";
import {feedlyConfig} from "../util/util";

export default async (request: VercelRequest, response: VercelResponse) => {
  const client = new AuthorizationCode(feedlyConfig);
  const refreshToken = request.query.refresh_token as string;

  try {
    let accessToken = client.createToken({refresh_token: refreshToken});
    accessToken = await accessToken.refresh();
    response.setHeader('Content-Type', 'application/json');
    response.end(JSON.stringify(accessToken));
  } catch (error) {
    console.log('Refresh Token Error', error.message);
    response.end();
  }
};




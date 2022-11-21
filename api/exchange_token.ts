import type {VercelRequest, VercelResponse} from "@vercel/node";
import {AuthorizationCode} from "simple-oauth2";
import {feedlyConfig, baseConfig} from "../util/util";

export default async (request: VercelRequest, response: VercelResponse) => {
  const client = new AuthorizationCode(feedlyConfig);

  try {
    const accessToken = await client.getToken({
      ...baseConfig,
      code: request.query.code as string,
    });
    response.setHeader('Content-Type', 'application/json');
    response.end(JSON.stringify(accessToken));

  } catch (error) {
    console.log('Access Token Error', error.message);
    response.end();
  }
};

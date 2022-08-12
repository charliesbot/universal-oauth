import type {VercelRequest, VercelResponse} from "@vercel/node";
import {AuthorizationCode} from "simple-oauth2";
import {feedlyConfig, baseConfig} from "../util/util";

export default async (_: VercelRequest, response: VercelResponse) => {
  const client = new AuthorizationCode(feedlyConfig);

  const authorizationUri = client.authorizeURL({
    ...baseConfig,
    state: 'pseudo-random',
  });

  console.log(authorizationUri)
  response.status(301).redirect(authorizationUri);
};



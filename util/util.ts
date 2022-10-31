const feedlyUrl = "https://cloud.feedly.com";
const clientID = process.env.CLIENT_ID;
const clientSecret = process.env.CLIENT_SECRET;

const baseConfig = {
  redirect_uri:
    "charliesbotrssapp://oauth",
  scope: 'https://cloud.feedly.com/subscriptions',
}

const feedlyConfig = {
  client: {
    id: clientID,
    secret: clientSecret,
  },
  auth: {
    authorizeHost: `${feedlyUrl}/v3/auth/auth`,
    authorizePath: `${feedlyUrl}/v3/auth/auth`,
    tokenHost: `${feedlyUrl}/v3/auth/token`,
    tokenPath: `${feedlyUrl}/v3/auth/token`,
  }
};

export {feedlyConfig, baseConfig};

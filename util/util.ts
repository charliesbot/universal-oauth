const feedlyUrl = "https://sandbox7.feedly.com";
const clientID = "sandbox";
const clientSecret = "NdyYvHssp6H6c2iTiU6mMaBQQ409pMOy";

const baseConfig = {
  redirect_uri:
    "http://localhost:8080/callback",
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
  }
};

export {feedlyConfig, baseConfig};

import type {VercelRequest, VercelResponse} from "@vercel/node";

export default (_: VercelRequest, response: VercelResponse) => {
  console.log("hola")
  response.status(200).send("Success Login");
};

import { NextRequestWithUnkeyContext, unstable__withUnkey } from "@unkey/nextjs";
import { NextResponse } from "next/server";

async function handler(req: NextRequestWithUnkeyContext) {
  if (!req.unkey.valid) {
    return new NextResponse("unauthorized", { status: 403 });
  }

  return new NextResponse(`Your API key is valid!

${JSON.stringify(req.unkey, null, 2)}`);
}

export const GET = unstable__withUnkey(handler, {
  /**
   * Just for this demo we're passing the key in searchparams,
   * by default it is loaded from the `Authorization` header
   */
  getKey: (req) => new URL(req.url).searchParams.get("key"),
});

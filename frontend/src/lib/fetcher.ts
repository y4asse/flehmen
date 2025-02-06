async function fetchGet<T>(
  url: string,
  token?: string,
  headers: Record<string, string> = {}
): Promise<T> {
  return fetchWrapper<T>(url, "GET", token, undefined, headers);
}

async function fetchPost<T, BodyType>(
  url: string,
  body: BodyType,
  token?: string,
  headers: Record<string, string> = {}
): Promise<T> {
  return fetchWrapper<T, BodyType>(url, "POST", token, body, headers);
}

async function fetchPut<T, BodyType>(
  url: string,
  body: BodyType,
  token?: string,
  headers: Record<string, string> = {}
): Promise<T> {
  return fetchWrapper<T, BodyType>(url, "PUT", token, body, headers);
}

async function fetchDelete<T>(
  url: string,
  token?: string,
  headers: Record<string, string> = {}
): Promise<T> {
  return fetchWrapper<T>(url, "DELETE", token, undefined, headers);
}

async function fetchWrapper<ResponseType, BodyType = undefined>(
  url: string,
  method: "GET" | "POST" | "PUT" | "DELETE",
  token?: string,
  body?: BodyType,
  headers: Record<string, string> = {}
): Promise<ResponseType> {
  if (!(body instanceof FormData) && !headers["Content-Type"]) {
    headers["Content-Type"] = "application/json";
  }

  if (token) {
    headers["Authorization"] = `Bearer ${token}`;
  }

  const options: RequestInit = {
    method,
    headers,
  };

  if (body !== undefined && !(body instanceof FormData)) {
    options.body = JSON.stringify(body);
  } else if (body instanceof FormData) {
    // body が FormData のインスタンスの場合は、そのまま設定
    options.body = body;
  }

  if (method === "GET" || method === "DELETE") {
    delete options.body;
  }

  try {
    const response = await fetch(
      `${process.env.NEXT_PUBLIC_API_URL}${url}`,
      options
    );

    if (!response.ok) {
      throw new Error(`Error: status code ${response.status}`);
    }

    return response.json() as Promise<ResponseType>;
  } catch (e) {
    console.log(e);
    return {
      status: 500,
      message: "Error: something went wrong",
    } as ResponseType;
  }
}

export { fetchGet, fetchPost, fetchPut, fetchDelete };

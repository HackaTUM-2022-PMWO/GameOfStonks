export class AuthorizationError extends Error {}
export class TransportError extends Error {}

export interface ServiceConstructor<ST> extends Function {
  defaultEndpoint: string;
  prototype: ST;

  new (transport: <T>(method: string, data?: any[]) => Promise<T>): ST;
}

export const getClient = <T>(clientClass: ServiceConstructor<T>) => {
  const transport = getTransport(clientClass.defaultEndpoint);
  return new clientClass(transport);
};

// utility for creating async transport clients
export const getTransport =
  (endpoint: string) =>
  async <T>(method: string, args: any[] = []) => {
    return new Promise<T>((resolve, reject) =>
      fetch(`${endpoint}/${encodeURIComponent(method)}`, {
        method: "POST",
        body: JSON.stringify(args),
      })
        .then((response) => {
          if (response.status !== 200) {
            if (response.status === 401) {
              throw new AuthorizationError(
                `Authorization required ${response.statusText}`
              );
            }
            throw new TransportError(response.statusText);
          }
          return response.json() as Promise<T>;
        })
        .then((val) => resolve(val))
        .catch((err) => reject(err))
    );
  };

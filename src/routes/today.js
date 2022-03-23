/**
 * [WIP] API endpoint to return today's answers
 * @returns today's date in ISO format.
 */
export async function today() {
  const now = new Date().toISOString();

  return {
    body: { now },
  };
}

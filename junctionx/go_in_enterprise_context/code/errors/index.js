try {
  const value = await queryDatabase();
  try {
    const response = await makeApiCall(value);
    console.log(response);
  } catch (err) {
    console.log(err.message);
  }
} catch (err) {
  console.log(err.message);
}

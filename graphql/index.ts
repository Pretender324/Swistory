import { ApolloClient, HttpLink, InMemoryCache } from '@apollo/client'

const link = new HttpLink({
  uri: 'http://localhost:4000/graphql',
  // Additional options
})

export const client = new ApolloClient({
  //TODO: productionに修正
  link,
  cache: new InMemoryCache(),
})

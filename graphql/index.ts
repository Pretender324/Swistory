import { ApolloClient, InMemoryCache } from '@apollo/client'

export const client = new ApolloClient({
  //TODO: productionに修正
  uri: 'http://localhost:4000',
  cache: new InMemoryCache(),
})

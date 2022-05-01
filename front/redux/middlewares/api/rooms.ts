import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { IBody, IRoomResponse, IRoomsResponse } from '@/lib/types'

// Define a service using a base URL and expected endpoints
export const roomsApi = createApi({
  reducerPath: 'roomsApi',
  baseQuery: fetchBaseQuery({ baseUrl: process.env.NEXT_PUBLIC_API }),
  endpoints: (builder) => ({
    getRooms: builder.query<IRoomsResponse, void>({
      query: () => `/rooms`,
      transformResponse: (resp: IBody) => resp.Value
    }),
    getRoom: builder.query<IRoomResponse, string>({
      query: (id: string) => `/rooms/${id}`
    })
  }),
})

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const { useGetRoomsQuery, useGetRoomQuery } = roomsApi
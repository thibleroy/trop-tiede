import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { IBody, IDeviceResponse, IDevicesResponse } from '@/lib/types'

// Define a service using a base URL and expected endpoints
export const devicesApi = createApi({
  reducerPath: 'devicesApi',
  baseQuery: fetchBaseQuery({ baseUrl: process.env.NEXT_PUBLIC_API }),
  endpoints: (builder) => ({
    getDevices: builder.query<IDevicesResponse, void>({
      query: () => `/devices`,
      transformResponse: (resp: IBody) => resp.Value
    }),
    getDevice: builder.query<IDeviceResponse, string>({
      query: (id: string) => `/devices/${id}`
    })
  }),
})

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const { useGetDeviceQuery, useGetDevicesQuery } = devicesApi

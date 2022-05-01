import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react'
import { IDevice, IDevicesResponse } from '@/lib/types'

// Define a service using a base URL and expected endpoints
export const devicesApi = createApi({
  reducerPath: 'devicesApi',
  baseQuery: fetchBaseQuery({ baseUrl: process.env.NEXT_PUBLIC_API }),
  endpoints: (builder) => ({
    getDevices: builder.query<IDevice[], string>({
      query: () => `/devices`,
      transformResponse: (resp: IDevicesResponse) => resp.Devices
    }),
    getDevice: builder.query<IDevice, string>({
      query: (id) => `/devices/${id}`
    })
  }),
})

// Export hooks for usage in functional components, which are
// auto-generated based on the defined endpoints
export const { useGetDeviceQuery, useGetDevicesQuery } = devicesApi
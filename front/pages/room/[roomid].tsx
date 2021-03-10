import { useRouter } from 'next/router'

const Room = () => {
    const router = useRouter()
    const { roomid } = router.query
    return <p>Post: {roomid}</p>
}

export default Room

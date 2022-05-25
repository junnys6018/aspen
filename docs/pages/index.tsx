import type { NextPage } from 'next';
import Head from 'next/head';
import Link from 'next/link';
import Footer from '../components/footer';
import Header from '../components/header';

const Home: NextPage = () => {
    return (
        <div>
            <Head>
                <title>Aspen</title>
                <link rel="icon" href="/favicon.ico" />
            </Head>
            <Header />
            <div className="container relative mx-auto my-8">
                <p>
                    This page is currently under construction, why dont you checkout our{' '}
                    <Link href="/playground">
                        <a className="text-blue-600 hover:underline">playgound</a>
                    </Link>{' '}
                    in the mean time?
                </p>
            </div>
            <Footer />
        </div>
    );
};

export default Home;

import Head from 'next/head';
import Image from 'next/image';
import { faEye, faRocket } from '@fortawesome/free-solid-svg-icons';
import { PreviewButton, GenerateButton } from '../components';
import { ToastContainer } from 'react-toastify';

export default function Home() {
	const element = (
		<div className="max-w-xl mx-auto mt-20 p-10 justify-center items-center">
			<Head>StockX GIF Generator</Head>

			<div className="mb-5 w-full flex justify-center">
				<Image src="/banner.png" height={230} width={600} className="shadow-sm" loading="eager" />
			</div>
			<h2 className="text-center font-bold text-4x1" style={{ fontWeight: 900 }}>
				StockX Product GIF Generator
			</h2>
			<h5 className="text-center">
				Create GIFs from StockX products with{' '}
				<a
					className="text-blue-400 hover:text-blue-600"
					href="https://stockx.com/news/360-degree-images-are-live/"
					target="_blank"
					rel="noreferrer"
				>
					360º image support
				</a>
				.
			</h5>
			<br />
			<br />
			<input
				className="h-9 w-full px-3 rounded-md placeholder-gray-500 placeholder-opacity-75 text-gray-800 outline-none focus:outline-none focus:ring transition-all focus:ring-blue-500"
				type="text"
				id="url"
				placeholder="StockX Link/Search"
			/>
			<div className="items-center justify-center pt-5 space-x-3 w-full flex">
				<PreviewButton icon={faEye} content={'Generate Preview'} />
				<GenerateButton icon={faRocket} content={'Generate GIF'} />
			</div>
			<div className="p-5 rounded-md bg-gradient-to-tr from-gray-800 to-gray-700 space-y-4 shadow-md mt-10">
				<p className="text-sm">Developer&apos;s Note:</p>
				<p className="text-sm">
					This tool&apos;s mobile functionalty is lacking and only consistently performs on iOS in Safari.
				</p>
				<p className="text-sm">
					It may take a while to download the file considering they normally run upwards of 7MB.
				</p>
			</div>
			<hr className="opacity-10 my-5" />
			<p className="text-center text-xs opacity-40">
				Made by{' '}
				<a
					className="text-blue-400 hover:text-blue-600"
					href="https://twitter.com/fykowo"
					target="_blank"
					rel="noreferrer"
				>
					Fyko
				</a>{' '}
				~ This service is not affiliated with nor endorsed by StockX LLC.
			</p>

			<ToastContainer
				position="bottom-right"
				autoClose={3000}
				hideProgressBar={false}
				newestOnTop={false}
				closeOnClick
				rtl={false}
				pauseOnFocusLoss
				draggable
				pauseOnHover
			/>
		</div>
	);

	return element;
}

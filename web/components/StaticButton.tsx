/* eslint-disable no-alert */
import { IconProp } from '@fortawesome/fontawesome-svg-core';
import { faSpinner } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useState } from 'react';
import { toast } from 'react-toastify';
import { fetchStatic, queryAlgolia } from '../util';

export function StaticButton({ icon: _icon, content: _content }: { icon: IconProp; content: string }) {
	const [disabled, setDisabled] = useState(false);
	const [icon, setIcon] = useState(_icon);
	const [spin, setSpin] = useState(false);
	const [content, setContent] = useState(_content);

	const cleanup = () => {
		setIcon(_icon);
		setSpin(false);
		setContent(_content);
		setDisabled(false);
	};

	const onClick = async () => {
		setDisabled(true);

		setContent('Fetching static URL...');
		setIcon(faSpinner);
		setSpin(true);

		const search = document.getElementById('url') as HTMLInputElement;
		const test = /stockx.com\/(?<slug>.*)/.exec(search.value);
		if (!test || !test.groups?.slug) {
			cleanup();
			return toast.error('No StockX product URL provided!');
		}

		const query = await queryAlgolia(test.groups.slug);
		if (!query.length) {
			cleanup();
			return toast.error('Product not found!');
		}

		const hit = query[0];
		const exists = await fetchStatic(hit.id, false, true);

		if (exists.status === 404) {
			setContent('Generating static URL...');
			const location = await fetchStatic(hit.id, false, false);
			const out = await location.text();
			alert(out);
		} else {
			const out = await exists.text();
			alert(out);
		}

		cleanup();
		return toast.success('Successfully fetched static GIF URL!');
	};

	return (
		<button onClick={onClick} className="bluebutton" disabled={disabled}>
			<FontAwesomeIcon icon={icon} spin={spin} /> {content}
		</button>
	);
}

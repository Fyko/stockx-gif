import { IconProp } from '@fortawesome/fontawesome-svg-core';
import { faSpinner } from '@fortawesome/free-solid-svg-icons';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { useState } from 'react';
import { toast } from 'react-toastify';
import { generateGif, queryAlgolia } from '../util';

export function PreviewButton({ icon: _icon, content: _content }: { icon: IconProp; content: string }) {
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

		setContent('Generating preview...');
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
		const gif = await generateGif(hit.id, true);
		if (gif.status === 409) {
			cleanup();
			return toast.error('The product does not have 360 image support!');
		}

		const url = URL.createObjectURL(await gif.blob());
		const a = document.createElement('a');
		a.href = url;
		a.download = `${hit.slug}-preview.gif`;
		document.body.appendChild(a);
		a.click();
		a.remove();
		URL.revokeObjectURL(url);

		cleanup();
		return toast.success('Successfully generated and downloaded GIF preview!');
	};

	return (
		<button onClick={onClick} className="bluebutton" disabled={disabled}>
			<FontAwesomeIcon icon={icon} spin={spin} /> {content}
		</button>
	);
}

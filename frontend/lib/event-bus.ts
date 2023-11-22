class EventBus {
	static instance: EventBus;

	private listeners: { [key: string]: Function[] } = {};

	static getInstance(): EventBus {
		if (!EventBus.instance) {
			EventBus.instance = new EventBus();
		}

		return EventBus.instance;
	}

	subscribe(event: string, callback: Function): void {
		if (!this.listeners[event]) {
			this.listeners[event] = [];
		}

		this.listeners[event].push(callback);
	}

	publish(event: string, ...args: any[]): void {
		if (this.listeners[event]) {
			this.listeners[event].forEach(callback => {
				callback(...args);
			});
		}
	}

	unsubscribe(event: string, callback: Function): void {
		if (this.listeners[event]) {
			this.listeners[event] = this.listeners[event].filter(listener => listener !== callback);
		}
	}

	unsubscribeAll(event: string): void {
		if (this.listeners[event]) {
			this.listeners[event] = [];
		}
	}
}

const eventBus = EventBus.getInstance();

export { eventBus };

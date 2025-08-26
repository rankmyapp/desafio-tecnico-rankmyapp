import { Processor, WorkerHost } from '@nestjs/bullmq';
import { Job } from 'bullmq';

@Processor('process-orders')
export class ProcessOrderConsumer extends WorkerHost {
  async process(job: Job<any, any, string>): Promise<any> {
    console.log('Processing job:', job.name, job.data);
    
    if (job.name === 'buy-ticket') {
      // const progress = 0;
      // for (let i = 0; i < 100; i++) {
        console.log('CHEGOU A MSG', job.data);
        // progress += 1;
        // await job.progress(progress);
      // }
    }
    
    return {};
  }
}
